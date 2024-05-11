/*
Copyright 2024.

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

// ApeCloudMySQLSpec defines the desired state of ApeCloudMySQL
type ApeCloudMySQLSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ApeCloudMySQL. Edit apecloudmysql_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ApeCloudMySQLStatus defines the observed state of ApeCloudMySQL
type ApeCloudMySQLStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ApeCloudMySQL is the Schema for the apecloudmysqls API
type ApeCloudMySQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApeCloudMySQLSpec   `json:"spec,omitempty"`
	Status ApeCloudMySQLStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApeCloudMySQLList contains a list of ApeCloudMySQL
type ApeCloudMySQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApeCloudMySQL `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApeCloudMySQL{}, &ApeCloudMySQLList{})
}