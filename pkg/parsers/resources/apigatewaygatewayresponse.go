package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayGatewayResponse Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html
type ApiGatewayGatewayResponse struct {
	Type       string                              `yaml:"Type"`
	Properties ApiGatewayGatewayResponseProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// ApiGatewayGatewayResponse Properties
type ApiGatewayGatewayResponseProperties struct {
	ResponseType       interface{} `yaml:"ResponseType"`
	RestApiId          interface{} `yaml:"RestApiId"`
	StatusCode         interface{} `yaml:"StatusCode,omitempty"`
	ResponseParameters interface{} `yaml:"ResponseParameters,omitempty"`
	ResponseTemplates  interface{} `yaml:"ResponseTemplates,omitempty"`
}

// NewApiGatewayGatewayResponse constructor creates a new ApiGatewayGatewayResponse
func NewApiGatewayGatewayResponse(properties ApiGatewayGatewayResponseProperties, deps ...interface{}) ApiGatewayGatewayResponse {
	return ApiGatewayGatewayResponse{
		Type:       "AWS::ApiGateway::GatewayResponse",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayGatewayResponse parses ApiGatewayGatewayResponse
func ParseApiGatewayGatewayResponse(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayGatewayResponse
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayGatewayResponse - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ApiGatewayGatewayResponse) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayGatewayResponseProperties) Validate() []error {
	errs := []error{}
	if resource.ResponseType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResponseType'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
