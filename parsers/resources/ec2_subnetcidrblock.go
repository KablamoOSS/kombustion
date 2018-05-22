package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2SubnetCidrBlock struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SubnetCidrBlockProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SubnetCidrBlockProperties struct {
	Ipv6CidrBlock interface{} `yaml:"Ipv6CidrBlock"`
	SubnetId interface{} `yaml:"SubnetId"`
}

func NewEC2SubnetCidrBlock(properties EC2SubnetCidrBlockProperties, deps ...interface{}) EC2SubnetCidrBlock {
	return EC2SubnetCidrBlock{
		Type:       "AWS::EC2::SubnetCidrBlock",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SubnetCidrBlock(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2SubnetCidrBlock
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SubnetCidrBlock - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2SubnetCidrBlock) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SubnetCidrBlockProperties) Validate() []error {
	errs := []error{}
	if resource.Ipv6CidrBlock == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Ipv6CidrBlock'"))
	}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errs
}
