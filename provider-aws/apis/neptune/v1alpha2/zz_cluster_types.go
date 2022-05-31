/*
Copyright 2022 Upbound Inc.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ClusterMembers []*string `json:"clusterMembers,omitempty" tf:"cluster_members,omitempty"`

	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ReaderEndpoint *string `json:"readerEndpoint,omitempty" tf:"reader_endpoint,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty" tf:"enable_cloudwatch_logs_exports,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IAMRoleIdRefs []v1.Reference `json:"iamRoleIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IAMRoleIdSelector *v1.Selector `json:"iamRoleIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:refFieldName=IAMRoleIdRefs
	// +crossplane:generate:reference:selectorFieldName=IAMRoleIdSelector
	// +kubebuilder:validation:Optional
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=ClusterParameterGroup
	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupName *string `json:"neptuneClusterParameterGroupName,omitempty" tf:"neptune_cluster_parameter_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupNameRef *v1.Reference `json:"neptuneClusterParameterGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupNameSelector *v1.Selector `json:"neptuneClusterParameterGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupNameRef *v1.Reference `json:"neptuneSubnetGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupNameSelector *v1.Selector `json:"neptuneSubnetGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifierRef *v1.Reference `json:"replicationSourceIdentifierRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifierSelector *v1.Selector `json:"replicationSourceIdentifierSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// +crossplane:generate:reference:type=ClusterSnapshot
	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotIdentifierRef *v1.Reference `json:"snapshotIdentifierRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SnapshotIdentifierSelector *v1.Selector `json:"snapshotIdentifierSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}