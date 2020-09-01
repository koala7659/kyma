package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *AuthorizationPolicySpec `json:"spec"`
}

type AuthorizationPolicySpec struct {
	Selector *WorkloadSelector `json:"selector,omitempty"`
	Action   Action            `json:"action,omitempty"`
	Rules    []Rule            `json:"rules,omitempty"`
}

type WorkloadSelector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}

type Rule struct {
	From []From      `json:"from,omitempty"`
	To   []To        `json:"to,omitempty"`
	When []Condition `json:"when,omitempty"`
}

type Source struct {
	Principals           []string `json:"principals,omitempty"`
	NotPrincipals        []string `json:"notPrincipals,omitempty"`
	RequestPrincipals    []string `json:"requestPrincipals,omitempty"`
	NotRequestPrincipals []string `json:"notRequestPrincipals,omitempty"`
	Namespaces           []string `json:"namespaces,omitempty"`
	NotNamespaces        []string `json:"notNamespaces,omitempty"`
	IpBlocks             []string `json:"ipBlocks,omitempty"`
	NotIpBlocks          []string `json:"notIpBlocks,omitempty"`
}

type Operation struct {
	Hosts      []string `json:"hosts,omitempty"`
	NotHosts   []string `json:"notHosts,omitempty"`
	Ports      []string `json:"ports,omitempty"`
	NotPorts   []string `json:"notPorts,omitempty"`
	Methods    []string `json:"methods,omitempty"`
	NotMethods []string `json:"notMethods,omitempty"`
	Paths      []string `json:"paths,omitempty"`
	NotPaths   []string `json:"notPaths,omitempty"`
}

type Condition struct {
	Key       string   `json:"key"`
	Values    []string `json:"values,omitempty"`
	NotValues []string `json:"notValues,omitempty"`
}

type From struct {
	Source Source `json:"source"`
}

type To struct {
	Operation Operation `json:"operation"`
}

type Action string

const (
	Allow Action = "ALLOW"
	Deny  Action = "DENY"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AuthorizationPolicy `json:"items"`
}

// TODO: Delete it!
// RuleSpec defines specification for Rule
type RuleSpec struct {
	Match   string       `json:"match"`
	Actions []RuleAction `json:"actions"`
}

// RuleAction defines action for Rule
type RuleAction struct {
	Handler   string   `json:"handler"`
	Instances []string `json:"instances"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RuleList is a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Rule `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Handler defines Istio Handler CR with Params for "denier" adapter
type Handler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *HandlerSpec `json:"spec"`
}

// HandlerSpec defines specification for Handler
type HandlerSpec struct {
	CompiledAdapter string `json:"compiledAdapter"`
	// Params have a different form for different adapters
	Params *DenierHandlerParams `json:"params"`
}

// DenierHandlerParams defines handler params for "denier" adapter
type DenierHandlerParams struct {
	Status *DenierStatus `json:"status"`
}

// DenierStatus defines status for Denier
type DenierStatus struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HandlerList is a list of Handlers
type HandlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Handler `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Instance defines Istio Instance CR
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *InstanceSpec `json:"spec"`
}

// InstanceSpec defines specification for Instance
type InstanceSpec struct {
	CompiledTemplate string `json:"compiledTemplate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstanceList is a list of Handlers
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Instance `json:"items"`
}
