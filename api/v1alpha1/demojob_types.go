/*

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

// DemoJobSpec defines the desired state of DemoJob
type DemoJobSpec struct {
	Image string `json:"image,omitempty"`
}

// DemoJobStatus defines the observed state of DemoJob
type DemoJobStatus struct {
}

// +kubebuilder:object:root=true

// DemoJob is the Schema for the demojobs API
type DemoJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoJobSpec   `json:"spec,omitempty"`
	Status DemoJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DemoJobList contains a list of DemoJob
type DemoJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoJob{}, &DemoJobList{})
}
