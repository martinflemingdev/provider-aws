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

type BrokerLogsObservation struct {
}

type BrokerLogsParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogs []CloudwatchLogsParameters `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs,omitempty"`

	// +kubebuilder:validation:Optional
	Firehose []FirehoseParameters `json:"firehose,omitempty" tf:"firehose,omitempty"`

	// +kubebuilder:validation:Optional
	S3 []S3Parameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type BrokerNodeGroupInfoObservation struct {
}

type BrokerNodeGroupInfoParameters struct {

	// The distribution of broker nodes across availability zones (documentation). Currently the only valid value is DEFAULT.
	// +kubebuilder:validation:Optional
	AzDistribution *string `json:"azDistribution,omitempty" tf:"az_distribution,omitempty"`

	// A list of subnets to connect to in client VPC (documentation).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	ClientSubnets []*string `json:"clientSubnets,omitempty" tf:"client_subnets,omitempty"`

	// References to Subnet in ec2 to populate clientSubnets.
	// +kubebuilder:validation:Optional
	ClientSubnetsRefs []v1.Reference `json:"clientSubnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate clientSubnets.
	// +kubebuilder:validation:Optional
	ClientSubnetsSelector *v1.Selector `json:"clientSubnetsSelector,omitempty" tf:"-"`

	// Information about the cluster access configuration. See below. For security reasons, you can't turn on public access while creating an MSK cluster. However, you can update an existing cluster to make it publicly accessible. You can also create a new cluster and then update it to make it publicly accessible (documentation).
	// +kubebuilder:validation:Optional
	ConnectivityInfo []ConnectivityInfoParameters `json:"connectivityInfo,omitempty" tf:"connectivity_info,omitempty"`

	// The size in GiB of the EBS volume for the data drive on each broker node.
	// +kubebuilder:validation:Optional
	EBSVolumeSize *float64 `json:"ebsVolumeSize,omitempty" tf:"ebs_volume_size,omitempty"`

	// Specify the instance type to use for the kafka brokersE.g., kafka.m5.large. (Pricing info)
	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsRefs []v1.Reference `json:"securityGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsSelector *v1.Selector `json:"securityGroupsSelector,omitempty" tf:"-"`

	// A block that contains information about storage volumes attached to MSK broker nodes. See below.
	// +kubebuilder:validation:Optional
	StorageInfo []StorageInfoParameters `json:"storageInfo,omitempty" tf:"storage_info,omitempty"`
}

type ClientAuthenticationObservation struct {
}

type ClientAuthenticationParameters struct {

	// Configuration block for specifying SASL client authentication. See below.
	// +kubebuilder:validation:Optional
	Sasl []SaslParameters `json:"sasl,omitempty" tf:"sasl,omitempty"`

	// Configuration block for specifying TLS client authentication. See below.
	// +kubebuilder:validation:Optional
	TLS []TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`

	// Enables unauthenticated access.
	// +kubebuilder:validation:Optional
	Unauthenticated *bool `json:"unauthenticated,omitempty" tf:"unauthenticated,omitempty"`
}

type CloudwatchLogsObservation struct {
}

type CloudwatchLogsParameters struct {

	// Controls whether provisioned throughput is enabled or not. Default value: false.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Name of the Cloudwatch Log Group to deliver logs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group
	// +kubebuilder:validation:Optional
	LogGroup *string `json:"logGroup,omitempty" tf:"log_group,omitempty"`

	// Reference to a Group in cloudwatchlogs to populate logGroup.
	// +kubebuilder:validation:Optional
	LogGroupRef *v1.Reference `json:"logGroupRef,omitempty" tf:"-"`

	// Selector for a Group in cloudwatchlogs to populate logGroup.
	// +kubebuilder:validation:Optional
	LogGroupSelector *v1.Selector `json:"logGroupSelector,omitempty" tf:"-"`
}

type ClusterObservation struct {

	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if encryption_info.0.encryption_in_transit.0.client_broker is set to PLAINTEXT or TLS_PLAINTEXT. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
	BootstrapBrokers *string `json:"bootstrapBrokers,omitempty" tf:"bootstrap_brokers,omitempty"`

	// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198. This attribute will have a value if encryption_info.0.encryption_in_transit.0.client_broker is set to TLS_PLAINTEXT or TLS and client_authentication.0.sasl.0.iam is set to true and broker_node_group_info.0.connectivity_info.0.public_access.0.type is set to SERVICE_PROVIDED_EIPS and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
	BootstrapBrokersPublicSaslIAM *string `json:"bootstrapBrokersPublicSaslIam,omitempty" tf:"bootstrap_brokers_public_sasl_iam,omitempty"`

	// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196. This attribute will have a value if encryption_info.0.encryption_in_transit.0.client_broker is set to TLS_PLAINTEXT or TLS and client_authentication.0.sasl.0.scram is set to true and broker_node_group_info.0.connectivity_info.0.public_access.0.type is set to SERVICE_PROVIDED_EIPS and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
	BootstrapBrokersPublicSaslScram *string `json:"bootstrapBrokersPublicSaslScram,omitempty" tf:"bootstrap_brokers_public_sasl_scram,omitempty"`

	// One or more DNS names (or IP addresses) and TLS port pairs. For example, b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194. This attribute will have a value if encryption_info.0.encryption_in_transit.0.client_broker is set to TLS_PLAINTEXT or TLS and broker_node_group_info.0.connectivity_info.0.public_access.0.type is set to SERVICE_PROVIDED_EIPS and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
	BootstrapBrokersPublicTLS *string `json:"bootstrapBrokersPublicTls,omitempty" tf:"bootstrap_brokers_public_tls,omitempty"`

	// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098. This attribute will have a value if encryption_info.0.encryption_in_transit.0.client_broker is set to TLS_PLAINTEXT or TLS and client_authentication.0.sasl.0.iam is set to true. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
	BootstrapBrokersSaslIAM *string `json:"bootstrapBrokersSaslIam,omitempty" tf:"bootstrap_brokers_sasl_iam,omitempty"`

	// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096. This attribute will have a value if encryption_info.0.encryption_in_transit.0.client_broker is set to TLS_PLAINTEXT or TLS and client_authentication.0.sasl.0.scram is set to true. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
	BootstrapBrokersSaslScram *string `json:"bootstrapBrokersSaslScram,omitempty" tf:"bootstrap_brokers_sasl_scram,omitempty"`

	// One or more DNS names (or IP addresses) and TLS port pairs. For example, b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094. This attribute will have a value if encryption_info.0.encryption_in_transit.0.client_broker is set to TLS_PLAINTEXT or TLS. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
	BootstrapBrokersTLS *string `json:"bootstrapBrokersTls,omitempty" tf:"bootstrap_brokers_tls,omitempty"`

	// Current version of the MSK Cluster used for updates, e.g., K13V1IB3VIYZZH
	CurrentVersion *string `json:"currentVersion,omitempty" tf:"current_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
	ZookeeperConnectString *string `json:"zookeeperConnectString,omitempty" tf:"zookeeper_connect_string,omitempty"`

	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster via TLS. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
	ZookeeperConnectStringTLS *string `json:"zookeeperConnectStringTls,omitempty" tf:"zookeeper_connect_string_tls,omitempty"`
}

type ClusterParameters struct {

	// Configuration block for the broker nodes of the Kafka cluster.
	// +kubebuilder:validation:Required
	BrokerNodeGroupInfo []BrokerNodeGroupInfoParameters `json:"brokerNodeGroupInfo" tf:"broker_node_group_info,omitempty"`

	// Configuration block for specifying a client authentication. See below.
	// +kubebuilder:validation:Optional
	ClientAuthentication []ClientAuthenticationParameters `json:"clientAuthentication,omitempty" tf:"client_authentication,omitempty"`

	// Name of the MSK cluster.
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	// +kubebuilder:validation:Optional
	ConfigurationInfo []ConfigurationInfoParameters `json:"configurationInfo,omitempty" tf:"configuration_info,omitempty"`

	// Configuration block for specifying encryption. See below.
	// +kubebuilder:validation:Optional
	EncryptionInfo []EncryptionInfoParameters `json:"encryptionInfo,omitempty" tf:"encryption_info,omitempty"`

	// Specify the desired enhanced MSK CloudWatch monitoring level. See Monitoring Amazon MSK with Amazon CloudWatch
	// +kubebuilder:validation:Optional
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty" tf:"enhanced_monitoring,omitempty"`

	// Specify the desired Kafka software version.
	// +kubebuilder:validation:Required
	KafkaVersion *string `json:"kafkaVersion" tf:"kafka_version,omitempty"`

	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	// +kubebuilder:validation:Optional
	LoggingInfo []LoggingInfoParameters `json:"loggingInfo,omitempty" tf:"logging_info,omitempty"`

	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	// +kubebuilder:validation:Required
	NumberOfBrokerNodes *float64 `json:"numberOfBrokerNodes" tf:"number_of_broker_nodes,omitempty"`

	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	// +kubebuilder:validation:Optional
	OpenMonitoring []OpenMonitoringParameters `json:"openMonitoring,omitempty" tf:"open_monitoring,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationInfoObservation struct {
}

type ConfigurationInfoParameters struct {

	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	// +kubebuilder:validation:Required
	Arn *string `json:"arn" tf:"arn,omitempty"`

	// Revision of the MSK Configuration to use in the cluster.
	// +kubebuilder:validation:Required
	Revision *float64 `json:"revision" tf:"revision,omitempty"`
}

type ConnectivityInfoObservation struct {
}

type ConnectivityInfoParameters struct {

	// Access control settings for brokers. See below.
	// +kubebuilder:validation:Optional
	PublicAccess []PublicAccessParameters `json:"publicAccess,omitempty" tf:"public_access,omitempty"`
}

type EBSStorageInfoObservation struct {
}

type EBSStorageInfoParameters struct {

	// A block that contains EBS volume provisioned throughput information. To provision storage throughput, you must choose broker type kafka.m5.4xlarge or larger. See below.
	// +kubebuilder:validation:Optional
	ProvisionedThroughput []ProvisionedThroughputParameters `json:"provisionedThroughput,omitempty" tf:"provisioned_throughput,omitempty"`

	// The size in GiB of the EBS volume for the data drive on each broker node. Minimum value of 1 and maximum value of 16384.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type EncryptionInTransitObservation struct {
}

type EncryptionInTransitParameters struct {

	// Encryption setting for data in transit between clients and brokers. Valid values: TLS, TLS_PLAINTEXT, and PLAINTEXT. Default value is TLS.
	// +kubebuilder:validation:Optional
	ClientBroker *string `json:"clientBroker,omitempty" tf:"client_broker,omitempty"`

	// Whether data communication among broker nodes is encrypted. Default value: true.
	// +kubebuilder:validation:Optional
	InCluster *bool `json:"inCluster,omitempty" tf:"in_cluster,omitempty"`
}

type EncryptionInfoObservation struct {
}

type EncryptionInfoParameters struct {

	// The ARN of the KMS key used for encryption at rest of the broker data volumes.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	EncryptionAtRestKMSKeyArn *string `json:"encryptionAtRestKmsKeyArn,omitempty" tf:"encryption_at_rest_kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate encryptionAtRestKmsKeyArn.
	// +kubebuilder:validation:Optional
	EncryptionAtRestKMSKeyArnRef *v1.Reference `json:"encryptionAtRestKmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate encryptionAtRestKmsKeyArn.
	// +kubebuilder:validation:Optional
	EncryptionAtRestKMSKeyArnSelector *v1.Selector `json:"encryptionAtRestKmsKeyArnSelector,omitempty" tf:"-"`

	// Configuration block to specify encryption in transit. See below.
	// +kubebuilder:validation:Optional
	EncryptionInTransit []EncryptionInTransitParameters `json:"encryptionInTransit,omitempty" tf:"encryption_in_transit,omitempty"`
}

type FirehoseObservation struct {
}

type FirehoseParameters struct {

	// Name of the Kinesis Data Firehose delivery stream to deliver logs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +kubebuilder:validation:Optional
	DeliveryStream *string `json:"deliveryStream,omitempty" tf:"delivery_stream,omitempty"`

	// Reference to a DeliveryStream in firehose to populate deliveryStream.
	// +kubebuilder:validation:Optional
	DeliveryStreamRef *v1.Reference `json:"deliveryStreamRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate deliveryStream.
	// +kubebuilder:validation:Optional
	DeliveryStreamSelector *v1.Selector `json:"deliveryStreamSelector,omitempty" tf:"-"`

	// Controls whether provisioned throughput is enabled or not. Default value: false.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type JmxExporterObservation struct {
}

type JmxExporterParameters struct {

	// Indicates whether you want to enable or disable the JMX Exporter.
	// +kubebuilder:validation:Required
	EnabledInBroker *bool `json:"enabledInBroker" tf:"enabled_in_broker,omitempty"`
}

type LoggingInfoObservation struct {
}

type LoggingInfoParameters struct {

	// Configuration block for Broker Logs settings for logging info. See below.
	// +kubebuilder:validation:Required
	BrokerLogs []BrokerLogsParameters `json:"brokerLogs" tf:"broker_logs,omitempty"`
}

type NodeExporterObservation struct {
}

type NodeExporterParameters struct {

	// Indicates whether you want to enable or disable the JMX Exporter.
	// +kubebuilder:validation:Required
	EnabledInBroker *bool `json:"enabledInBroker" tf:"enabled_in_broker,omitempty"`
}

type OpenMonitoringObservation struct {
}

type OpenMonitoringParameters struct {

	// Configuration block for Prometheus settings for open monitoring. See below.
	// +kubebuilder:validation:Required
	Prometheus []PrometheusParameters `json:"prometheus" tf:"prometheus,omitempty"`
}

type PrometheusObservation struct {
}

type PrometheusParameters struct {

	// Configuration block for JMX Exporter. See below.
	// +kubebuilder:validation:Optional
	JmxExporter []JmxExporterParameters `json:"jmxExporter,omitempty" tf:"jmx_exporter,omitempty"`

	// Configuration block for Node Exporter. See below.
	// +kubebuilder:validation:Optional
	NodeExporter []NodeExporterParameters `json:"nodeExporter,omitempty" tf:"node_exporter,omitempty"`
}

type ProvisionedThroughputObservation struct {
}

type ProvisionedThroughputParameters struct {

	// Controls whether provisioned throughput is enabled or not. Default value: false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second. The minimum value is 250. The maximum value varies between broker type. You can refer to the valid values for the maximum volume throughput at the following documentation on throughput bottlenecks
	// +kubebuilder:validation:Optional
	VolumeThroughput *float64 `json:"volumeThroughput,omitempty" tf:"volume_throughput,omitempty"`
}

type PublicAccessObservation struct {
}

type PublicAccessParameters struct {

	// Public access type. Valida values: DISABLED, SERVICE_PROVIDED_EIPS.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type S3Observation struct {
}

type S3Parameters struct {

	// Name of the S3 bucket to deliver logs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Controls whether provisioned throughput is enabled or not. Default value: false.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Prefix to append to the folder name.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SaslObservation struct {
}

type SaslParameters struct {

	// Enables IAM client authentication. Defaults to false.
	// +kubebuilder:validation:Optional
	IAM *bool `json:"iam,omitempty" tf:"iam,omitempty"`

	// Enables SCRAM client authentication via AWS Secrets Manager. Defaults to false.
	// +kubebuilder:validation:Optional
	Scram *bool `json:"scram,omitempty" tf:"scram,omitempty"`
}

type StorageInfoObservation struct {
}

type StorageInfoParameters struct {

	// A block that contains EBS volume information. See below.
	// +kubebuilder:validation:Optional
	EBSStorageInfo []EBSStorageInfoParameters `json:"ebsStorageInfo,omitempty" tf:"ebs_storage_info,omitempty"`
}

type TLSObservation struct {
}

type TLSParameters struct {

	// List of ACM Certificate Authority Amazon Resource Names (ARNs).
	// +kubebuilder:validation:Optional
	CertificateAuthorityArns []*string `json:"certificateAuthorityArns,omitempty" tf:"certificate_authority_arns,omitempty"`
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

// Cluster is the Schema for the Clusters API. Upbound official provider resource for managing an aws managed streaming for kafka cluster
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