package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DAXCluster struct {
	Type       string                      `yaml:"Type"`
	Properties DAXClusterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DAXClusterProperties struct {
	ClusterName interface{} `yaml:"ClusterName,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	IAMRoleARN interface{} `yaml:"IAMRoleARN"`
	NodeType interface{} `yaml:"NodeType"`
	NotificationTopicARN interface{} `yaml:"NotificationTopicARN,omitempty"`
	ParameterGroupName interface{} `yaml:"ParameterGroupName,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	ReplicationFactor interface{} `yaml:"ReplicationFactor"`
	SubnetGroupName interface{} `yaml:"SubnetGroupName,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	AvailabilityZones interface{} `yaml:"AvailabilityZones,omitempty"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
}

func NewDAXCluster(properties DAXClusterProperties, deps ...interface{}) DAXCluster {
	return DAXCluster{
		Type:       "AWS::DAX::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDAXCluster(name string, data string) (cf types.ValueMap, err error) {
	var resource DAXCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DAXCluster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DAXCluster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DAXClusterProperties) Validate() []error {
	errs := []error{}
	if resource.IAMRoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IAMRoleARN'"))
	}
	if resource.NodeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NodeType'"))
	}
	if resource.ReplicationFactor == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationFactor'"))
	}
	return errs
}
