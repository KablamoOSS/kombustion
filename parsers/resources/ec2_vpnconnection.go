package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPNConnection struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPNConnectionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPNConnectionProperties struct {
	CustomerGatewayId interface{} `yaml:"CustomerGatewayId"`
	StaticRoutesOnly interface{} `yaml:"StaticRoutesOnly,omitempty"`
	Type interface{} `yaml:"Type"`
	VpnGatewayId interface{} `yaml:"VpnGatewayId"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VpnTunnelOptionsSpecifications interface{} `yaml:"VpnTunnelOptionsSpecifications,omitempty"`
}

func NewEC2VPNConnection(properties EC2VPNConnectionProperties, deps ...interface{}) EC2VPNConnection {
	return EC2VPNConnection{
		Type:       "AWS::EC2::VPNConnection",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPNConnection(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPNConnection
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPNConnection - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPNConnection) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPNConnectionProperties) Validate() []error {
	errs := []error{}
	if resource.CustomerGatewayId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CustomerGatewayId'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.VpnGatewayId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpnGatewayId'"))
	}
	return errs
}
