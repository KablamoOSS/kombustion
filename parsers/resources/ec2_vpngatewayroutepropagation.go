package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPNGatewayRoutePropagation struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPNGatewayRoutePropagationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPNGatewayRoutePropagationProperties struct {
	VpnGatewayId interface{} `yaml:"VpnGatewayId"`
	RouteTableIds interface{} `yaml:"RouteTableIds"`
}

func NewEC2VPNGatewayRoutePropagation(properties EC2VPNGatewayRoutePropagationProperties, deps ...interface{}) EC2VPNGatewayRoutePropagation {
	return EC2VPNGatewayRoutePropagation{
		Type:       "AWS::EC2::VPNGatewayRoutePropagation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPNGatewayRoutePropagation(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPNGatewayRoutePropagation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPNGatewayRoutePropagation - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPNGatewayRoutePropagation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPNGatewayRoutePropagationProperties) Validate() []error {
	errs := []error{}
	if resource.VpnGatewayId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpnGatewayId'"))
	}
	if resource.RouteTableIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RouteTableIds'"))
	}
	return errs
}
