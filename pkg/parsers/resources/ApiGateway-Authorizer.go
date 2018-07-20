package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayAuthorizer Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html
type ApiGatewayAuthorizer struct {
	Type       string                         `yaml:"Type"`
	Properties ApiGatewayAuthorizerProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// ApiGatewayAuthorizer Properties
type ApiGatewayAuthorizerProperties struct {
	AuthType                     interface{} `yaml:"AuthType,omitempty"`
	AuthorizerCredentials        interface{} `yaml:"AuthorizerCredentials,omitempty"`
	AuthorizerResultTtlInSeconds interface{} `yaml:"AuthorizerResultTtlInSeconds,omitempty"`
	AuthorizerUri                interface{} `yaml:"AuthorizerUri,omitempty"`
	IdentitySource               interface{} `yaml:"IdentitySource,omitempty"`
	IdentityValidationExpression interface{} `yaml:"IdentityValidationExpression,omitempty"`
	Name                         interface{} `yaml:"Name,omitempty"`
	RestApiId                    interface{} `yaml:"RestApiId"`
	Type                         interface{} `yaml:"Type,omitempty"`
	ProviderARNs                 interface{} `yaml:"ProviderARNs,omitempty"`
}

// NewApiGatewayAuthorizer constructor creates a new ApiGatewayAuthorizer
func NewApiGatewayAuthorizer(properties ApiGatewayAuthorizerProperties, deps ...interface{}) ApiGatewayAuthorizer {
	return ApiGatewayAuthorizer{
		Type:       "AWS::ApiGateway::Authorizer",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayAuthorizer parses ApiGatewayAuthorizer
func ParseApiGatewayAuthorizer(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource ApiGatewayAuthorizer
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseApiGatewayAuthorizer validator
func (resource ApiGatewayAuthorizer) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayAuthorizerProperties validator
func (resource ApiGatewayAuthorizerProperties) Validate() []error {
	errors := []error{}
	if resource.RestApiId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errors
}
