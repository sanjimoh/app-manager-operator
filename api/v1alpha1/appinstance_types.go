/*
Copyright 2025.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppInstanceSpec defines the desired state of AppInstance
type AppInstanceSpec struct {
	// Size is the number of replicas of the application
	Size int32 `json:"size"`

	// Image is the container image for the application
	Image string `json:"image"`

	// Environment represents the deployment environment (dev, staging, prod)
	Environment string `json:"environment"`

	// Resources specifies the resource requirements for the application
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// AppInstanceStatus defines the observed state of AppInstance
type AppInstanceStatus struct {
	// Nodes are the names of the pods
	Nodes []string `json:"nodes"`

	// Status of the deployment
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AppInstance is the Schema for the appinstances API
type AppInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppInstanceSpec   `json:"spec,omitempty"`
	Status AppInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppInstanceList contains a list of AppInstance
type AppInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppInstance{}, &AppInstanceList{})
}
