package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BruslanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Bruslan `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Bruslan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              BruslanSpec   `json:"spec"`
	Status            BruslanStatus `json:"status,omitempty"`
}

type BruslanSpec struct {
	// Fill me
	BruslanDEMO     bool   `json ",bruslanDemoBool"`
	NameSpaceString string `json ",nameSpaceString"`
}
type BruslanStatus struct {
	// Fill me
	Status string `json ",status"`
}
