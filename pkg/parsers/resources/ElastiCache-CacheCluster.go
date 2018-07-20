package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ElastiCacheCacheCluster Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html
type ElastiCacheCacheCluster struct {
	Type       string                            `yaml:"Type"`
	Properties ElastiCacheCacheClusterProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// ElastiCacheCacheCluster Properties
type ElastiCacheCacheClusterProperties struct {
	AZMode                     interface{} `yaml:"AZMode,omitempty"`
	AutoMinorVersionUpgrade    interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	CacheNodeType              interface{} `yaml:"CacheNodeType"`
	CacheParameterGroupName    interface{} `yaml:"CacheParameterGroupName,omitempty"`
	CacheSubnetGroupName       interface{} `yaml:"CacheSubnetGroupName,omitempty"`
	ClusterName                interface{} `yaml:"ClusterName,omitempty"`
	Engine                     interface{} `yaml:"Engine"`
	EngineVersion              interface{} `yaml:"EngineVersion,omitempty"`
	NotificationTopicArn       interface{} `yaml:"NotificationTopicArn,omitempty"`
	NumCacheNodes              interface{} `yaml:"NumCacheNodes"`
	Port                       interface{} `yaml:"Port,omitempty"`
	PreferredAvailabilityZone  interface{} `yaml:"PreferredAvailabilityZone,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	SnapshotName               interface{} `yaml:"SnapshotName,omitempty"`
	SnapshotRetentionLimit     interface{} `yaml:"SnapshotRetentionLimit,omitempty"`
	SnapshotWindow             interface{} `yaml:"SnapshotWindow,omitempty"`
	CacheSecurityGroupNames    interface{} `yaml:"CacheSecurityGroupNames,omitempty"`
	PreferredAvailabilityZones interface{} `yaml:"PreferredAvailabilityZones,omitempty"`
	SnapshotArns               interface{} `yaml:"SnapshotArns,omitempty"`
	Tags                       interface{} `yaml:"Tags,omitempty"`
	VpcSecurityGroupIds        interface{} `yaml:"VpcSecurityGroupIds,omitempty"`
}

// NewElastiCacheCacheCluster constructor creates a new ElastiCacheCacheCluster
func NewElastiCacheCacheCluster(properties ElastiCacheCacheClusterProperties, deps ...interface{}) ElastiCacheCacheCluster {
	return ElastiCacheCacheCluster{
		Type:       "AWS::ElastiCache::CacheCluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElastiCacheCacheCluster parses ElastiCacheCacheCluster
func ParseElastiCacheCacheCluster(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource ElastiCacheCacheCluster
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseElastiCacheCacheCluster validator
func (resource ElastiCacheCacheCluster) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElastiCacheCacheClusterProperties validator
func (resource ElastiCacheCacheClusterProperties) Validate() []error {
	errors := []error{}
	if resource.CacheNodeType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'CacheNodeType'"))
	}
	if resource.Engine == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Engine'"))
	}
	if resource.NumCacheNodes == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'NumCacheNodes'"))
	}
	return errors
}
