package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EC2NatGateway struct {
	Type       string                  `yaml:"Type"`
	Properties EC2NatGatewayProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

type EC2NatGatewayProperties struct {
	AllocationId interface{} `yaml:"AllocationId"`
	SubnetId     interface{} `yaml:"SubnetId"`
	Tags         interface{} `yaml:"Tags,omitempty"`
}

func NewEC2NatGateway(properties EC2NatGatewayProperties, deps ...interface{}) EC2NatGateway {
	return EC2NatGateway{
		Type:       "AWS::EC2::NatGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2NatGateway(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2NatGateway
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2NatGateway - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2NatGateway) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2NatGatewayProperties) Validate() []error {
	errs := []error{}
	if resource.AllocationId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AllocationId'"))
	}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errs
}
