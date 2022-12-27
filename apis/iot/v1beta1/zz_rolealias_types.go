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

type RoleAliasObservation struct {

	// The ARN assigned by AWS to this role alias.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleAliasParameters struct {

	// The name of the role alias.
	// +kubebuilder:validation:Required
	Alias *string `json:"alias" tf:"alias,omitempty"`

	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	// +kubebuilder:validation:Optional
	CredentialDuration *float64 `json:"credentialDuration,omitempty" tf:"credential_duration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The identity of the role to which the alias refers.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// RoleAliasSpec defines the desired state of RoleAlias
type RoleAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAliasParameters `json:"forProvider"`
}

// RoleAliasStatus defines the observed state of RoleAlias.
type RoleAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAlias is the Schema for the RoleAliass API. Provides an IoT role alias.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RoleAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAliasSpec   `json:"spec"`
	Status            RoleAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAliasList contains a list of RoleAliass
type RoleAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAlias `json:"items"`
}

// Repository type metadata.
var (
	RoleAlias_Kind             = "RoleAlias"
	RoleAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAlias_Kind}.String()
	RoleAlias_KindAPIVersion   = RoleAlias_Kind + "." + CRDGroupVersion.String()
	RoleAlias_GroupVersionKind = CRDGroupVersion.WithKind(RoleAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAlias{}, &RoleAliasList{})
}