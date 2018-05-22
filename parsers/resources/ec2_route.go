package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2Route struct {
	Type       string                      `yaml:"Type"`
	Properties EC2RouteProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2RouteProperties struct {
	DestinationCidrBlock interface{} `yaml:"DestinationCidrBlock,omitempty"`
	DestinationIpv6CidrBlock interface{} `yaml:"DestinationIpv6CidrBlock,omitempty"`
	EgressOnlyInternetGatewayId interface{} `yaml:"EgressOnlyInternetGatewayId,omitempty"`
	GatewayId interface{} `yaml:"GatewayId,omitempty"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
	NatGatewayId interface{} `yaml:"NatGatewayId,omitempty"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId,omitempty"`
	RouteTableId interface{} `yaml:"RouteTableId"`
	VpcPeeringConnectionId interface{} `yaml:"VpcPeeringConnectionId,omitempty"`
}

func NewEC2Route(properties EC2RouteProperties, deps ...interface{}) EC2Route {
	return EC2Route{
		Type:       "AWS::EC2::Route",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2Route(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2Route
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2Route - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2Route) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2RouteProperties) Validate() []error {
	errs := []error{}
	if resource.RouteTableId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RouteTableId'"))
	}
	return errs
}
