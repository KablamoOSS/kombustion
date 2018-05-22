package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPNConnectionRoute struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPNConnectionRouteProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPNConnectionRouteProperties struct {
	DestinationCidrBlock interface{} `yaml:"DestinationCidrBlock"`
	VpnConnectionId interface{} `yaml:"VpnConnectionId"`
}

func NewEC2VPNConnectionRoute(properties EC2VPNConnectionRouteProperties, deps ...interface{}) EC2VPNConnectionRoute {
	return EC2VPNConnectionRoute{
		Type:       "AWS::EC2::VPNConnectionRoute",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPNConnectionRoute(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPNConnectionRoute
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPNConnectionRoute - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPNConnectionRoute) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPNConnectionRouteProperties) Validate() []error {
	errs := []error{}
	if resource.DestinationCidrBlock == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DestinationCidrBlock'"))
	}
	if resource.VpnConnectionId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpnConnectionId'"))
	}
	return errs
}
