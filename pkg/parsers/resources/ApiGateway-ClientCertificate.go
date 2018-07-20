package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayClientCertificate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html
type ApiGatewayClientCertificate struct {
	Type       string                                `yaml:"Type"`
	Properties ApiGatewayClientCertificateProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// ApiGatewayClientCertificate Properties
type ApiGatewayClientCertificateProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
}

// NewApiGatewayClientCertificate constructor creates a new ApiGatewayClientCertificate
func NewApiGatewayClientCertificate(properties ApiGatewayClientCertificateProperties, deps ...interface{}) ApiGatewayClientCertificate {
	return ApiGatewayClientCertificate{
		Type:       "AWS::ApiGateway::ClientCertificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayClientCertificate parses ApiGatewayClientCertificate
func ParseApiGatewayClientCertificate(
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
	var resource ApiGatewayClientCertificate
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

// ParseApiGatewayClientCertificate validator
func (resource ApiGatewayClientCertificate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayClientCertificateProperties validator
func (resource ApiGatewayClientCertificateProperties) Validate() []error {
	errors := []error{}
	return errors
}
