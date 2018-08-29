package v1alpha1

import meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SslConfig struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               SslConfigSpec   `json:"spec"`
	Status             SslConfigStatus `json:"status,omitempty"`
}
type SslConfigSpec struct {
	Cert   string `json:"cert"`
	Key    string `json:"key"`
	Domain string `json:"domain"`
}

type SslConfigStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

type SslConfigList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []SslConfig `json:"items"`
}
