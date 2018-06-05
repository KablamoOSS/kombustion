package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type ApiGatewayMethod struct {
	Type       string                     `yaml:"Type"`
	Properties ApiGatewayMethodProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

type ApiGatewayMethodProperties struct {
	ApiKeyRequired     interface{}                    `yaml:"ApiKeyRequired,omitempty"`
	AuthorizationType  interface{}                    `yaml:"AuthorizationType,omitempty"`
	AuthorizerId       interface{}                    `yaml:"AuthorizerId,omitempty"`
	HttpMethod         interface{}                    `yaml:"HttpMethod"`
	OperationName      interface{}                    `yaml:"OperationName,omitempty"`
	RequestValidatorId interface{}                    `yaml:"RequestValidatorId,omitempty"`
	ResourceId         interface{}                    `yaml:"ResourceId"`
	RestApiId          interface{}                    `yaml:"RestApiId"`
	RequestModels      interface{}                    `yaml:"RequestModels,omitempty"`
	RequestParameters  interface{}                    `yaml:"RequestParameters,omitempty"`
	MethodResponses    interface{}                    `yaml:"MethodResponses,omitempty"`
	Integration        *properties.Method_Integration `yaml:"Integration,omitempty"`
}

func NewApiGatewayMethod(properties ApiGatewayMethodProperties, deps ...interface{}) ApiGatewayMethod {
	return ApiGatewayMethod{
		Type:       "AWS::ApiGateway::Method",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayMethod(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayMethod
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayMethod - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ApiGatewayMethod) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayMethodProperties) Validate() []error {
	errs := []error{}
	if resource.HttpMethod == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HttpMethod'"))
	}
	if resource.ResourceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceId'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
