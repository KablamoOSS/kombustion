package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VPCEndpoint struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VPCEndpointProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VPCEndpointProperties struct {
	PolicyDocument interface{} `yaml:"PolicyDocument,omitempty"`
	ServiceName interface{} `yaml:"ServiceName"`
	VpcId interface{} `yaml:"VpcId"`
	RouteTableIds interface{} `yaml:"RouteTableIds,omitempty"`
}

func NewEC2VPCEndpoint(properties EC2VPCEndpointProperties, deps ...interface{}) EC2VPCEndpoint {
	return EC2VPCEndpoint{
		Type:       "AWS::EC2::VPCEndpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPCEndpoint(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VPCEndpoint
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCEndpoint - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VPCEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPCEndpointProperties) Validate() []error {
	errs := []error{}
	if resource.ServiceName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceName'"))
	}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
