package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2NetworkInterfacePermission struct {
	Type       string                      `yaml:"Type"`
	Properties EC2NetworkInterfacePermissionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2NetworkInterfacePermissionProperties struct {
	AwsAccountId interface{} `yaml:"AwsAccountId"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId"`
	Permission interface{} `yaml:"Permission"`
}

func NewEC2NetworkInterfacePermission(properties EC2NetworkInterfacePermissionProperties, deps ...interface{}) EC2NetworkInterfacePermission {
	return EC2NetworkInterfacePermission{
		Type:       "AWS::EC2::NetworkInterfacePermission",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2NetworkInterfacePermission(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2NetworkInterfacePermission
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2NetworkInterfacePermission - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2NetworkInterfacePermission) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2NetworkInterfacePermissionProperties) Validate() []error {
	errs := []error{}
	if resource.AwsAccountId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AwsAccountId'"))
	}
	if resource.NetworkInterfaceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NetworkInterfaceId'"))
	}
	if resource.Permission == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Permission'"))
	}
	return errs
}
