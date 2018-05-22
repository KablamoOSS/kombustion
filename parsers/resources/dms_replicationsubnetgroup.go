package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DMSReplicationSubnetGroup struct {
	Type       string                      `yaml:"Type"`
	Properties DMSReplicationSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DMSReplicationSubnetGroupProperties struct {
	ReplicationSubnetGroupDescription interface{} `yaml:"ReplicationSubnetGroupDescription"`
	ReplicationSubnetGroupIdentifier interface{} `yaml:"ReplicationSubnetGroupIdentifier,omitempty"`
	SubnetIds interface{} `yaml:"SubnetIds"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewDMSReplicationSubnetGroup(properties DMSReplicationSubnetGroupProperties, deps ...interface{}) DMSReplicationSubnetGroup {
	return DMSReplicationSubnetGroup{
		Type:       "AWS::DMS::ReplicationSubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDMSReplicationSubnetGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource DMSReplicationSubnetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSReplicationSubnetGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DMSReplicationSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSReplicationSubnetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.ReplicationSubnetGroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationSubnetGroupDescription'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
