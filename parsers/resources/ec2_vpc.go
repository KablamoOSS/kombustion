package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPC struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPCProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPCProperties struct {
	CidrBlock interface{} `yaml:"CidrBlock"`
	EnableDnsHostnames interface{} `yaml:"EnableDnsHostnames,omitempty"`
	EnableDnsSupport interface{} `yaml:"EnableDnsSupport,omitempty"`
	InstanceTenancy interface{} `yaml:"InstanceTenancy,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2VPC(properties EC2VPCProperties, deps ...interface{}) EC2VPC {
	return EC2VPC{
		Type:       "AWS::EC2::VPC",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPC(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPC
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPC - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPC) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPCProperties) Validate() []error {
	errs := []error{}
	if resource.CidrBlock == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CidrBlock'"))
	}
	return errs
}
