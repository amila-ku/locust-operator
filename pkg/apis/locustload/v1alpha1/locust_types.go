package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LocustSpec defines the desired state of Locust
type LocustSpec struct {
	//HostURL is the url the loadtest is executed agains
	HostURL string `json:"hosturl"`
	//Image is the container with locust files tests
	Image string `json:"image"`
	//NumberOfUsers is the maximum number of users to simulate
	Workers int32 `json:"workers,omitempty"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// LocustStatus defines the observed state of Locust
type LocustStatus struct {
	CurrentWorkers int32 `json:"currentworkers,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Locust is the Schema for the locusts API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=locusts,scope=Namespaced
type Locust struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocustSpec   `json:"spec,omitempty"`
	Status LocustStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LocustList contains a list of Locust
type LocustList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Locust `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Locust{}, &LocustList{})
}
