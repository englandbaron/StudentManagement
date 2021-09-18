package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Student struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StudentSpec   `json:"spec"`
	StudentStatus     StudentStatus `json:"status"`
}

// CicdJobSpec is the status for a CicdJob resource
type StudentSpec struct {
	Name string `json:"name"`
}

// CicdJobStatus is the status for a CicdJob resource
type StudentStatus struct {
	Eating   bool `json:"eating"`
	Learning bool `json:"learning"`
	Playing  bool `json:"playing"`
	Sleeping bool `json:"playing"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StudentList is a list of Student resources
type StudentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Student `json:"items"`
}
