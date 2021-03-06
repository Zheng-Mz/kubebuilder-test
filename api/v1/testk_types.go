/*
Copyright 2020 zmz.

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

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TestKSpec defines the desired state of TestK
type TestKSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TestK. Edit TestK_types.go to remove/update
	Foo string `json:"foo,omitempty"`

	//ServicePort
	ServicePort []corev1.ServicePort `json:"servicePort,omitempty"`

	// Specifies the deployment that will be created when executing a TestK.
	DeploySpec appsv1.DeploymentSpec `json:"deploySpec"`
}

// TestKStatus defines the observed state of TestK
type TestKStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// TestK is the Schema for the testks API
type TestK struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestKSpec   `json:"spec,omitempty"`
	Status TestKStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TestKList contains a list of TestK
type TestKList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestK `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestK{}, &TestKList{})
}
