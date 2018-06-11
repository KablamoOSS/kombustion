package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// DAXSubnetGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html
type DAXSubnetGroup struct {
	Type       string                   `yaml:"Type"`
	Properties DAXSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// DAXSubnetGroup Properties
type DAXSubnetGroupProperties struct {
	Description     interface{} `yaml:"Description,omitempty"`
	SubnetGroupName interface{} `yaml:"SubnetGroupName,omitempty"`
	SubnetIds       interface{} `yaml:"SubnetIds"`
}

// NewDAXSubnetGroup constructor creates a new DAXSubnetGroup
func NewDAXSubnetGroup(properties DAXSubnetGroupProperties, deps ...interface{}) DAXSubnetGroup {
	return DAXSubnetGroup{
		Type:       "AWS::DAX::SubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDAXSubnetGroup parses DAXSubnetGroup
func ParseDAXSubnetGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
