package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2SubnetNetworkAclAssociation struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SubnetNetworkAclAssociationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SubnetNetworkAclAssociationProperties struct {
	NetworkAclId interface{} `yaml:"NetworkAclId"`
	SubnetId interface{} `yaml:"SubnetId"`
}

func NewEC2SubnetNetworkAclAssociation(properties EC2SubnetNetworkAclAssociationProperties, deps ...interface{}) EC2SubnetNetworkAclAssociation {
	return EC2SubnetNetworkAclAssociation{
		Type:       "AWS::EC2::SubnetNetworkAclAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SubnetNetworkAclAssociation(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2SubnetNetworkAclAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SubnetNetworkAclAssociation - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2SubnetNetworkAclAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SubnetNetworkAclAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.NetworkAclId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NetworkAclId'"))
	}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errs
}
