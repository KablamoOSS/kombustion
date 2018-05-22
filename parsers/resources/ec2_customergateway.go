package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2CustomerGateway struct {
	Type       string                      `yaml:"Type"`
	Properties EC2CustomerGatewayProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2CustomerGatewayProperties struct {
	BgpAsn interface{} `yaml:"BgpAsn"`
	IpAddress interface{} `yaml:"IpAddress"`
	Type interface{} `yaml:"Type"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2CustomerGateway(properties EC2CustomerGatewayProperties, deps ...interface{}) EC2CustomerGateway {
	return EC2CustomerGateway{
		Type:       "AWS::EC2::CustomerGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2CustomerGateway(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2CustomerGateway
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2CustomerGateway - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2CustomerGateway) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2CustomerGatewayProperties) Validate() []error {
	errs := []error{}
	if resource.BgpAsn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BgpAsn'"))
	}
	if resource.IpAddress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpAddress'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
