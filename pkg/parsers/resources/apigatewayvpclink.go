package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayVpcLink Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html
type ApiGatewayVpcLink struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayVpcLinkProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// ApiGatewayVpcLink Properties
type ApiGatewayVpcLinkProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name        interface{} `yaml:"Name"`
	TargetArns  interface{} `yaml:"TargetArns"`
}

// NewApiGatewayVpcLink constructor creates a new ApiGatewayVpcLink
func NewApiGatewayVpcLink(properties ApiGatewayVpcLinkProperties, deps ...interface{}) ApiGatewayVpcLink {
	return ApiGatewayVpcLink{
		Type:       "AWS::ApiGateway::VpcLink",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayVpcLink parses ApiGatewayVpcLink
func ParseApiGatewayVpcLink(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
