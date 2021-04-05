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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LocustSpec defines the desired state of Locust
type LocustSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//HostURL is the url the loadtest is executed agains
	HostURL string `json:"hosturl"`
	//Image is the container with locust files tests
	Image string `json:"image"`
	//Users is the maximum number of users to simulate
	Users int `json:"users,omitempty"`
	//HatchRate is the maximum number of users to simulate
	HatchRate int `json:"hatchrate,omitempty"`
	//Slaves is the number of worker instances
	Slaves int32 `json:"slaves,omitempty"` 
	//MaxSlaves is the number of maximum worker instances
	MaxSlaves int32 `json:"maxSlaves,omitempty"` 
}

// LocustStatus defines the observed state of Locust
type LocustStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CurrentWorkers int32 `json:"currentworkers,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Locust is the Schema for the locusts API
type Locust struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocustSpec   `json:"spec,omitempty"`
	Status LocustStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LocustList contains a list of Locust
type LocustList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Locust `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Locust{}, &LocustList{})
}
