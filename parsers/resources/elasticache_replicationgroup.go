package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElastiCacheReplicationGroup struct {
	Type       string                      `yaml:"Type"`
	Properties ElastiCacheReplicationGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElastiCacheReplicationGroupProperties struct {
	AtRestEncryptionEnabled interface{} `yaml:"AtRestEncryptionEnabled,omitempty"`
	AuthToken interface{} `yaml:"AuthToken,omitempty"`
	AutoMinorVersionUpgrade interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	AutomaticFailoverEnabled interface{} `yaml:"AutomaticFailoverEnabled,omitempty"`
	CacheNodeType interface{} `yaml:"CacheNodeType,omitempty"`
	CacheParameterGroupName interface{} `yaml:"CacheParameterGroupName,omitempty"`
	CacheSubnetGroupName interface{} `yaml:"CacheSubnetGroupName,omitempty"`
	Engine interface{} `yaml:"Engine,omitempty"`
	EngineVersion interface{} `yaml:"EngineVersion,omitempty"`
	NotificationTopicArn interface{} `yaml:"NotificationTopicArn,omitempty"`
	NumCacheClusters interface{} `yaml:"NumCacheClusters,omitempty"`
	NumNodeGroups interface{} `yaml:"NumNodeGroups,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	PrimaryClusterId interface{} `yaml:"PrimaryClusterId,omitempty"`
	ReplicasPerNodeGroup interface{} `yaml:"ReplicasPerNodeGroup,omitempty"`
	ReplicationGroupDescription interface{} `yaml:"ReplicationGroupDescription"`
	ReplicationGroupId interface{} `yaml:"ReplicationGroupId,omitempty"`
	SnapshotName interface{} `yaml:"SnapshotName,omitempty"`
	SnapshotRetentionLimit interface{} `yaml:"SnapshotRetentionLimit,omitempty"`
	SnapshotWindow interface{} `yaml:"SnapshotWindow,omitempty"`
	SnapshottingClusterId interface{} `yaml:"SnapshottingClusterId,omitempty"`
	TransitEncryptionEnabled interface{} `yaml:"TransitEncryptionEnabled,omitempty"`
	CacheSecurityGroupNames interface{} `yaml:"CacheSecurityGroupNames,omitempty"`
	NodeGroupConfiguration interface{} `yaml:"NodeGroupConfiguration,omitempty"`
	PreferredCacheClusterAZs interface{} `yaml:"PreferredCacheClusterAZs,omitempty"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
	SnapshotArns interface{} `yaml:"SnapshotArns,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewElastiCacheReplicationGroup(properties ElastiCacheReplicationGroupProperties, deps ...interface{}) ElastiCacheReplicationGroup {
	return ElastiCacheReplicationGroup{
		Type:       "AWS::ElastiCache::ReplicationGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElastiCacheReplicationGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource ElastiCacheReplicationGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheReplicationGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElastiCacheReplicationGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElastiCacheReplicationGroupProperties) Validate() []error {
	errs := []error{}
	if resource.ReplicationGroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationGroupDescription'"))
	}
	return errs
}
