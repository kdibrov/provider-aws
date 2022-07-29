/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayMulticastDomainObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TransitGatewayMulticastDomainParameters struct {

	// +kubebuilder:validation:Optional
	AutoAcceptSharedAssociations *string `json:"autoAcceptSharedAssociations,omitempty" tf:"auto_accept_shared_associations,omitempty"`

	// +kubebuilder:validation:Optional
	Igmpv2Support *string `json:"igmpv2Support,omitempty" tf:"igmpv2_support,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	StaticSourcesSupport *string `json:"staticSourcesSupport,omitempty" tf:"static_sources_support,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TransitGatewayID *string `json:"transitGatewayId" tf:"transit_gateway_id,omitempty"`
}

// TransitGatewayMulticastDomainSpec defines the desired state of TransitGatewayMulticastDomain
type TransitGatewayMulticastDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayMulticastDomainParameters `json:"forProvider"`
}

// TransitGatewayMulticastDomainStatus defines the observed state of TransitGatewayMulticastDomain.
type TransitGatewayMulticastDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayMulticastDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastDomain is the Schema for the TransitGatewayMulticastDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayMulticastDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayMulticastDomainSpec   `json:"spec"`
	Status            TransitGatewayMulticastDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastDomainList contains a list of TransitGatewayMulticastDomains
type TransitGatewayMulticastDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayMulticastDomain `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayMulticastDomain_Kind             = "TransitGatewayMulticastDomain"
	TransitGatewayMulticastDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayMulticastDomain_Kind}.String()
	TransitGatewayMulticastDomain_KindAPIVersion   = TransitGatewayMulticastDomain_Kind + "." + CRDGroupVersion.String()
	TransitGatewayMulticastDomain_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayMulticastDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayMulticastDomain{}, &TransitGatewayMulticastDomainList{})
}