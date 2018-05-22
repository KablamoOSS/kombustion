package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2EgressOnlyInternetGateway struct {
	Type       string                      `yaml:"Type"`
	Properties EC2EgressOnlyInternetGatewayProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2EgressOnlyInternetGatewayProperties struct {
	VpcId interface{} `yaml:"VpcId"`
}

func NewEC2EgressOnlyInternetGateway(properties EC2EgressOnlyInternetGatewayProperties, deps ...interface{}) EC2EgressOnlyInternetGateway {
	return EC2EgressOnlyInternetGateway{
		Type:       "AWS::EC2::EgressOnlyInternetGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2EgressOnlyInternetGateway(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2EgressOnlyInternetGateway
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2EgressOnlyInternetGateway - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2EgressOnlyInternetGateway) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2EgressOnlyInternetGatewayProperties) Validate() []error {
	errs := []error{}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
