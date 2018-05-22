package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPNGateway struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPNGatewayProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPNGatewayProperties struct {
	AmazonSideAsn interface{} `yaml:"AmazonSideAsn,omitempty"`
	Type interface{} `yaml:"Type"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2VPNGateway(properties EC2VPNGatewayProperties, deps ...interface{}) EC2VPNGateway {
	return EC2VPNGateway{
		Type:       "AWS::EC2::VPNGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPNGateway(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPNGateway
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPNGateway - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPNGateway) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPNGatewayProperties) Validate() []error {
	errs := []error{}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
