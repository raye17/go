/*
Copyright 2023.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UserSpec struct {
	Username    string   `json:"username,omitempty"`
	Password    string   `json:"password,omitempty"`
	SecretName  string   `json:"secretName,omitempty"`
	Enabled     bool     `json:"enabled,omitempty" norman:"default=true"`
	Namespace   string   `json:"namespace,omitempty"`
	Description string   `json:"description,omitempty"`
	AdminRole   bool     `json:"adminRole,omitempty" norman:"default=false,writeOnly,noupdate"`
	Roles       []string `json:"roles,omitempty"`
}

// UserStatus defines the observed state of User
type UserStatus struct {
	// Type of user condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time that password was updated.
	PasswordUpdateTime string `json:"passwordUpdateTime,omitempty"`
	//preserve reference auth files
	AuthFile map[string]string `json:"authFile,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message  string `json:"message,omitempty"`
	Password string `json:"password,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
