/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	ascalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	locustloadv1alpha1 "github.com/amila-ku/locust-operator/api/v1alpha1"
)

// LocustReconciler reconciles a Locust object
type LocustReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=locustload.cndev.io,resources=locusts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=locustload.cndev.io,resources=locusts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=locustload.cndev.io,resources=locusts/finalizers,verbs=update

//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods;service,verbs=get;list;
//+kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscaler,verbs=get;list;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Locust object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *LocustReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//_ = r.Log.WithValues("locust", req.NamespacedName)
	reqLogger := r.Log.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)
	reqLogger.Info("Reconciling Locust")

	// Fetch the Locust locustResource
	locustResource := &locustloadv1alpha1.Locust{}
	err := r.Get(context.TODO(), req.NamespacedName, locustResource)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	// Define a new Deployment object
	deployment := r.deploymentForLocust(locustResource)

	// Set Locust locustResource as the owner and controller
	if err := ctrl.SetControllerReference(locustResource, deployment, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Check if this Deployment already exists
	found := &appsv1.Deployment{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		err = r.Create(context.TODO(), deployment)
		if err != nil {
			return ctrl.Result{}, err
		}

		// Deployment created successfully - don't requeue
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// Deployment already exists - don't requeue
	reqLogger.Info("Skip reconcile: Deployment already exists", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

	// Service
	service := r.serviceForLocust(locustResource)

	// Set Locust locustResource as the owner and controller
	if err := ctrl.SetControllerReference(locustResource, service, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Check if this Service already exists
	foundsvc := &corev1.Service{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundsvc)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
		err = r.Create(context.TODO(), service)
		if err != nil {
			return ctrl.Result{}, err
		}

		// Service created successfully - don't requeue
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// Service already exists - don't requeue
	reqLogger.Info("Skip reconcile: Service already exists", "Service.Namespace", foundsvc.Namespace, "Service.Name", foundsvc.Name)

	// Locust worker deployment, limit for maximum number of slaves set to 30
	if locustResource.Spec.Slaves != 0 && locustResource.Spec.Slaves < 30 {
		slavedeployment := r.deploymentForLocustSlaves(locustResource)

		// Set Locust locustResource as the owner and controller
		if err := ctrl.SetControllerReference(locustResource, slavedeployment, r.Scheme); err != nil {
			return ctrl.Result{}, err
		}

		// Check if this Deployment already exists
		foundslaves := &appsv1.Deployment{}
		err = r.Get(context.TODO(), types.NamespacedName{Name: slavedeployment.Name, Namespace: slavedeployment.Namespace}, foundslaves)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Locust Worker Deployment", "Deployment.Namespace", slavedeployment.Namespace, "Deployment.Name", slavedeployment.Name)
			err = r.Create(context.TODO(), slavedeployment)
			if err != nil {
				return ctrl.Result{}, err
			}

			// Deployment created successfully - don't requeue
			return ctrl.Result{}, nil
		} else if err != nil {
			return ctrl.Result{}, err
		}

		// Deployment already exists - don't requeue
		reqLogger.Info("Skip reconcile: Locust Worker Deployment already exists", "Deployment.Namespace", foundslaves.Namespace, "Deployment.Name", foundslaves.Name)

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LocustReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&locustloadv1alpha1.Locust{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&ascalingv1.HorizontalPodAutoscaler{}).
		Complete(r)
}

// deploymentForLocust returns a Locust Deployment object
func (r *LocustReconciler) deploymentForLocust(cr *locustloadv1alpha1.Locust) *appsv1.Deployment {
	ls := labelsForLocust(cr.Name)
	replicas := int32Ptr(1)

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   cr.Spec.Image,
						Name:    cr.Name,
						Command: []string{"locust", "-f", "/tasks/main.py", "--master", "-H", cr.Spec.HostURL},
						Env: []corev1.EnvVar{
							{
								Name:  "TARGET_HOST",
								Value: cr.Spec.HostURL,
							},
						},
						Ports: []corev1.ContainerPort{
							{
								Name:          "http",
								Protocol:      corev1.ProtocolTCP,
								ContainerPort: 8089,
							},
							{
								Name:          "worker-1",
								Protocol:      corev1.ProtocolTCP,
								ContainerPort: 5557,
							},
							{
								Name:          "worker-2",
								Protocol:      corev1.ProtocolTCP,
								ContainerPort: 5558,
							},
						},
					}},
				},
			},
		},
	}
	// Set Locust locustResource as the owner and controller
	ctrl.SetControllerReference(cr, dep, r.Scheme)
	return dep
}

// deploymentForLocustSlaves returns a Locust Deployment object
func (r *LocustReconciler) deploymentForLocustSlaves(cr *locustloadv1alpha1.Locust) *appsv1.Deployment {
	ls := labelsForLocust(cr.Name + "-worker")

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-worker",
			Namespace: cr.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &cr.Spec.Slaves,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   cr.Spec.Image,
						Name:    cr.Name + "-worker",
						Command: []string{"locust", "--worker", "--master-host", cr.Name + "-service", "-f", "/tasks/main.py"},
					}},
				},
			},
		},
	}
	// Set Locust locustResource as the owner and controller
	ctrl.SetControllerReference(cr, dep, r.Scheme)
	return dep
}

// serviceForLocust returns a Service object
func (r *LocustReconciler) serviceForLocust(cr *locustloadv1alpha1.Locust) *corev1.Service {
	ls := labelsForLocust(cr.Name)

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-service",
			Namespace: cr.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: ls,
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Protocol: corev1.ProtocolTCP,
					Port:     8089,
				},
				{
					Name:     "worker-1",
					Protocol: corev1.ProtocolTCP,
					Port:     5557,
				},
				{
					Name:     "worker-2",
					Protocol: corev1.ProtocolTCP,
					Port:     5558,
				},
			},
		},
	}
	// Set Locust locustResource as the owner and controller
	ctrl.SetControllerReference(cr, svc, r.Scheme)
	return svc
}

// hpaForLocust creates a horizontal pod autoscaler in kubernetes for locust slaves
func (r *LocustReconciler) hpaForLocust(cr *locustloadv1alpha1.Locust) *ascalingv1.HorizontalPodAutoscaler {
	ls := labelsForLocust(cr.Name)
	targetCPUUtilization := int32Ptr(60)

	svc := &ascalingv1.HorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-service",
			Namespace: cr.Namespace,
			Labels:    ls,
		},
		Spec: ascalingv1.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: ascalingv1.CrossVersionObjectReference{
				Kind: "Deployment",
				Name: cr.Name + "-worker",
			},
			MinReplicas:                    &cr.Spec.Slaves,
			MaxReplicas:                    cr.Spec.MaxSlaves,
			TargetCPUUtilizationPercentage: targetCPUUtilization,
		},
	}
	// Set Locust locustResource as the owner and controller
	ctrl.SetControllerReference(cr, svc, r.Scheme)
	return svc
}

// labelsForLocust returns the labels for selecting the resources
// belonging to the given Locust CR name.
func labelsForLocust(name string) map[string]string {
	return map[string]string{"app": "Locust", "Locust_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}

func int32Ptr(i int32) *int32 { return &i }
