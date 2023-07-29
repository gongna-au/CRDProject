package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLInstanceSpec defines the desired state of MySQLInstance
type MySQLInstanceSpec struct {
	// Username for the MySQL instance
	Username string `json:"username,omitempty"`

	// Password for the MySQL instance
	Password string `json:"password,omitempty"`

	// Database is the name of the default database
	Database string `json:"database,omitempty"`

	// Version of the MySQL instance
	Version string `json:"version,omitempty"`
}

// MySQLInstanceStatus defines the observed state of MySQLInstance
type MySQLInstanceStatus struct {
	// Conditions represent the latest available observations of an object's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The Connection Port of the MySQL instance
	ConnectionPort int `json:"connectionPort,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MySQLInstance is the Schema for the mysqlinstances API
type MySQLInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLInstanceSpec   `json:"spec,omitempty"`
	Status MySQLInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MySQLInstanceList contains a list of MySQLInstance
type MySQLInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLInstance{}, &MySQLInstanceList{})
}
