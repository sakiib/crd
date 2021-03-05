/*
Copyright 2015 The Kubernetes Authors.
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

// +groupName=book.com
// +kubebuilder:object:generator=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=bookapis,singular=bookapi,shortName=book,categories={}
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:printcolumn:JSONPath=".status.deploymentName",name=Deployment,type=string
// +kubebuilder:printcolumn:JSONPath=".status.serviceName",name=Service,type=string
// +kubebuilder:printcolumn:JSONPath=".status.replicas",name=Replicas,type=integer
// +kubebuilder:printcolumn:JSONPath=".status.phase",name=Status,type=string
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BookAPI is a top-level type. A client is created for it.
type BookAPI struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BookAPISpec `json:"spec"`
	// +optional
	Status BookAPIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BookAPIList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type BookAPIList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BookAPI `json:"items"`
}

// +kubebuilder:validation:Enum=ClusterIP;NodePort
type ServiceType string

type BookAPISpec struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	Replica *int64 `json:"replica"`
	Image   string `json:"image"`

	//+optional
	Port int64 `json:"port,omitempty"`
	//+optional
	Username string `json:"username,omitempty"`
	//+optional
	Password string `json:"password,omitempty"`
	// +optional
	ServiceName string `json:"serviceName"`

	// +kubebuilder:default=ClusterIP
	ServiceType ServiceType `json:"serviceType"`
	NodePort    int32       `json:"nodePort,omitempty"`
}

type BookAPIStatus struct {
	Phase        string `json:"phase,omitempty"`
	ReplicaCount int64  `json:"replica_count,omitempty"`
}
