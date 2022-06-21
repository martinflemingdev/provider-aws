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

type AbortIncompleteMultipartUploadObservation struct {
}

type AbortIncompleteMultipartUploadParameters struct {

	// +kubebuilder:validation:Optional
	DaysAfterInitiation *float64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AndObservation struct {
}

type AndParameters struct {

	// +kubebuilder:validation:Optional
	ObjectSizeGreaterThan *float64 `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectSizeLessThan *float64 `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketLifecycleConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketLifecycleConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Rule []BucketLifecycleConfigurationRuleParameters `json:"rule" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationRuleObservation struct {
}

type BucketLifecycleConfigurationRuleParameters struct {

	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// +kubebuilder:validation:Optional
	Expiration []RuleExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []RuleFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []RuleNoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []RuleNoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Transition []RuleTransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type RuleExpirationObservation struct {
}

type RuleExpirationParameters struct {

	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type RuleFilterObservation struct {
}

type RuleFilterParameters struct {

	// +kubebuilder:validation:Optional
	And []AndParameters `json:"and,omitempty" tf:"and,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectSizeGreaterThan *string `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectSizeLessThan *string `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleNoncurrentVersionExpirationObservation struct {
}

type RuleNoncurrentVersionExpirationParameters struct {

	// +kubebuilder:validation:Optional
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// +kubebuilder:validation:Optional
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type RuleNoncurrentVersionTransitionObservation struct {
}

type RuleNoncurrentVersionTransitionParameters struct {

	// +kubebuilder:validation:Optional
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// +kubebuilder:validation:Optional
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`

	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type RuleTransitionObservation struct {
}

type RuleTransitionParameters struct {

	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type TagObservation struct {
}

type TagParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// BucketLifecycleConfigurationSpec defines the desired state of BucketLifecycleConfiguration
type BucketLifecycleConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLifecycleConfigurationParameters `json:"forProvider"`
}

// BucketLifecycleConfigurationStatus defines the observed state of BucketLifecycleConfiguration.
type BucketLifecycleConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfiguration is the Schema for the BucketLifecycleConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketLifecycleConfigurationSpec   `json:"spec"`
	Status            BucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfigurationList contains a list of BucketLifecycleConfigurations
type BucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketLifecycleConfiguration_Kind             = "BucketLifecycleConfiguration"
	BucketLifecycleConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLifecycleConfiguration_Kind}.String()
	BucketLifecycleConfiguration_KindAPIVersion   = BucketLifecycleConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketLifecycleConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketLifecycleConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLifecycleConfiguration{}, &BucketLifecycleConfigurationList{})
}
