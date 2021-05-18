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
	"reflect"
	"testing"

	"github.com/amila-ku/locust-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	//"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	k8sreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var logger = logf.Log.WithName("unit-tests")

func TestLocustReconciler_Reconcile(t *testing.T) {
	nsn := types.NamespacedName{Name: "my-instance", Namespace: "default"}
	created := &v1alpha1.Locust{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nsn.Name,
			Namespace: nsn.Namespace,
		},
		Spec: v1alpha1.LocustSpec{
			HostURL: "https://test.com",
			Image:   "amilaku/locust:v0.0.2",
			Users:   2,
		},
	}

	// Add the Locust as a resource type
	testScheme.AddKnownTypes(v1alpha1.SchemeBuilder.GroupVersion, created)

	// Objects to track in the fake client.
	objs := []runtime.Object{created}

	// Create a fake client to mock API calls.
	cl := fake.NewFakeClient(objs...)

	//err := k8sClient.Create(context.Background(), created)

	req := k8sreconcile.Request{
		NamespacedName: nsn,
	}

	type fields struct {
		Client client.Client
		Log    logr.Logger
		Scheme *runtime.Scheme
	}
	type args struct {
		ctx context.Context
		req ctrl.Request
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ctrl.Result
		wantErr bool
	}{
		{
			name:    "TestNewObjects",
			fields:  fields{cl, logger, testScheme},
			args:    args{context.Background(), req},
			want:    ctrl.Result{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &LocustReconciler{
				Client: tt.fields.Client,
				Log:    tt.fields.Log,
				Scheme: tt.fields.Scheme,
			}
			got, err := r.Reconcile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocustReconciler.Reconcile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocustReconciler.Reconcile() = %v, want %v", got, tt.want)
			}
		})
	}
}
