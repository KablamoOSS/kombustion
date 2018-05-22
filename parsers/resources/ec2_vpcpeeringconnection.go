package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPCPeeringConnection struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPCPeeringConnectionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPCPeeringConnectionProperties struct {
	PeerOwnerId interface{} `yaml:"PeerOwnerId,omitempty"`
	PeerRoleArn interface{} `yaml:"PeerRoleArn,omitempty"`
	PeerVpcId interface{} `yaml:"PeerVpcId"`
	VpcId interface{} `yaml:"VpcId"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2VPCPeeringConnection(properties EC2VPCPeeringConnectionProperties, deps ...interface{}) EC2VPCPeeringConnection {
	return EC2VPCPeeringConnection{
		Type:       "AWS::EC2::VPCPeeringConnection",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPCPeeringConnection(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPCPeeringConnection
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCPeeringConnection - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPCPeeringConnection) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPCPeeringConnectionProperties) Validate() []error {
	errs := []error{}
	if resource.PeerVpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PeerVpcId'"))
	}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
