package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RedshiftClusterSubnetGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RedshiftClusterSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RedshiftClusterSubnetGroupProperties struct {
	Description interface{} `yaml:"Description"`
	SubnetIds interface{} `yaml:"SubnetIds"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRedshiftClusterSubnetGroup(properties RedshiftClusterSubnetGroupProperties, deps ...interface{}) RedshiftClusterSubnetGroup {
	return RedshiftClusterSubnetGroup{
		Type:       "AWS::Redshift::ClusterSubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRedshiftClusterSubnetGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RedshiftClusterSubnetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftClusterSubnetGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RedshiftClusterSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RedshiftClusterSubnetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
