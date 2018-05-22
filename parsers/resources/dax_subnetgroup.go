package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DAXSubnetGroup struct {
	Type       string                      `yaml:"Type"`
	Properties DAXSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DAXSubnetGroupProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	SubnetGroupName interface{} `yaml:"SubnetGroupName,omitempty"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

func NewDAXSubnetGroup(properties DAXSubnetGroupProperties, deps ...interface{}) DAXSubnetGroup {
	return DAXSubnetGroup{
		Type:       "AWS::DAX::SubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDAXSubnetGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource DAXSubnetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DAXSubnetGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DAXSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DAXSubnetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
