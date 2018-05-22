package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElastiCacheCacheCluster struct {
	Type       string                      `yaml:"Type"`
	Properties ElastiCacheCacheClusterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElastiCacheCacheClusterProperties struct {
	AZMode interface{} `yaml:"AZMode,omitempty"`
	AutoMinorVersionUpgrade interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	CacheNodeType interface{} `yaml:"CacheNodeType"`
	CacheParameterGroupName interface{} `yaml:"CacheParameterGroupName,omitempty"`
	CacheSubnetGroupName interface{} `yaml:"CacheSubnetGroupName,omitempty"`
	ClusterName interface{} `yaml:"ClusterName,omitempty"`
	Engine interface{} `yaml:"Engine"`
	EngineVersion interface{} `yaml:"EngineVersion,omitempty"`
	NotificationTopicArn interface{} `yaml:"NotificationTopicArn,omitempty"`
	NumCacheNodes interface{} `yaml:"NumCacheNodes"`
	Port interface{} `yaml:"Port,omitempty"`
	PreferredAvailabilityZone interface{} `yaml:"PreferredAvailabilityZone,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	SnapshotName interface{} `yaml:"SnapshotName,omitempty"`
	SnapshotRetentionLimit interface{} `yaml:"SnapshotRetentionLimit,omitempty"`
	SnapshotWindow interface{} `yaml:"SnapshotWindow,omitempty"`
	CacheSecurityGroupNames interface{} `yaml:"CacheSecurityGroupNames,omitempty"`
	PreferredAvailabilityZones interface{} `yaml:"PreferredAvailabilityZones,omitempty"`
	SnapshotArns interface{} `yaml:"SnapshotArns,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VpcSecurityGroupIds interface{} `yaml:"VpcSecurityGroupIds,omitempty"`
}

func NewElastiCacheCacheCluster(properties ElastiCacheCacheClusterProperties, deps ...interface{}) ElastiCacheCacheCluster {
	return ElastiCacheCacheCluster{
		Type:       "AWS::ElastiCache::CacheCluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElastiCacheCacheCluster(name string, data string) (cf types.ValueMap, err error) {
	var resource ElastiCacheCacheCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheCacheCluster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElastiCacheCacheCluster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElastiCacheCacheClusterProperties) Validate() []error {
	errs := []error{}
	if resource.CacheNodeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CacheNodeType'"))
	}
	if resource.Engine == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Engine'"))
	}
	if resource.NumCacheNodes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NumCacheNodes'"))
	}
	return errs
}
