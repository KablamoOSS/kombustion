package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayVpcLink struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayVpcLinkProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayVpcLinkProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	TargetArns interface{} `yaml:"TargetArns"`
}

func NewApiGatewayVpcLink(properties ApiGatewayVpcLinkProperties, deps ...interface{}) ApiGatewayVpcLink {
	return ApiGatewayVpcLink{
		Type:       "AWS::ApiGateway::VpcLink",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayVpcLink(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayVpcLink
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayVpcLink - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayVpcLink) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayVpcLinkProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.TargetArns == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetArns'"))
	}
	return errs
}
