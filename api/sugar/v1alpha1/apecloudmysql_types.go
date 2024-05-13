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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApeCloudMySQLSpec defines the desired state of ApeCloudMySQL
type ApeCloudMySQLSpec struct {
	BaseSpec `json:",inline"`

	// Specified the MySQL Component Spec.
	MySQLComponentSpec ClusterComponentSpec `json:"mySQLComponentSpec"`

	// Specified the Proxy Component Spec.
	//
	// +optional
	ProxyComponentSpec *ClusterComponentSpec `json:"proxyComponentSpec,omitempty"`
}

// ApeCloudMySQLStatus defines the observed state of ApeCloudMySQL
type ApeCloudMySQLStatus struct {
	BaseStatus `json:",inline"`
}

type ApeCloudMySQLTopology string

func (in *ApeCloudMySQLSpec) TranslateTo() *appsv1alpha1.ClusterSpec {
	clusterSpec := (&in.BaseSpec).TranslateTo()
	mysqlSpec := (&in.MySQLComponentSpec).TranslateTo()
	//mysqlSpec.ComponentDef = "mysql"
	mysqlSpec.ComponentDefRef = "mysql"
	clusterSpec.ComponentSpecs = append(clusterSpec.ComponentSpecs, *mysqlSpec)
	if in.ProxyComponentSpec != nil {
		proxySpec := in.ProxyComponentSpec.TranslateTo()
		vtGateSpec := proxySpec.DeepCopy()
		vtGateSpec.Name = fmt.Sprintf("%s-%s", proxySpec.Name, "vtgate")
		vtGateSpec.ComponentDef = "vtgate"
		vtControllerSpec := proxySpec.DeepCopy()
		vtControllerSpec.Name = fmt.Sprintf("%s-%s", proxySpec.Name, "vtcontroller")
		vtControllerSpec.ComponentDef = "vtcontroller"
		clusterSpec.ComponentSpecs = append(clusterSpec.ComponentSpecs, *vtGateSpec, *vtControllerSpec)
	}
	clusterSpec.ClusterDefRef = "apecloud-mysql"
	clusterSpec.ClusterVersionRef = "ac-mysql-8.0.30"
	return clusterSpec
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:resource:categories={kubeblocks,all}
// +kubebuilder:printcolumn:name="TERMINATION-POLICY",type="string",JSONPath=".spec.terminationPolicy",description="Cluster termination policy."
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.phase",description="Cluster Status."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"

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
