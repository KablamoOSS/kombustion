package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2Subnet struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SubnetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SubnetProperties struct {
	AssignIpv6AddressOnCreation interface{} `yaml:"AssignIpv6AddressOnCreation,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	CidrBlock interface{} `yaml:"CidrBlock"`
	Ipv6CidrBlock interface{} `yaml:"Ipv6CidrBlock,omitempty"`
	MapPublicIpOnLaunch interface{} `yaml:"MapPublicIpOnLaunch,omitempty"`
	VpcId interface{} `yaml:"VpcId"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2Subnet(properties EC2SubnetProperties, deps ...interface{}) EC2Subnet {
	return EC2Subnet{
		Type:       "AWS::EC2::Subnet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2Subnet(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2Subnet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2Subnet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2Subnet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SubnetProperties) Validate() []error {
	errs := []error{}
	if resource.CidrBlock == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CidrBlock'"))
	}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
