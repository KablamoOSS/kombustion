package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayApiKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
type ApiGatewayApiKey struct {
	Type       string                     `yaml:"Type"`
	Properties ApiGatewayApiKeyProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// ApiGatewayApiKey Properties
type ApiGatewayApiKeyProperties struct {
	CustomerId         interface{} `yaml:"CustomerId,omitempty"`
	Description        interface{} `yaml:"Description,omitempty"`
	Enabled            interface{} `yaml:"Enabled,omitempty"`
	GenerateDistinctId interface{} `yaml:"GenerateDistinctId,omitempty"`
	Name               interface{} `yaml:"Name,omitempty"`
	StageKeys          interface{} `yaml:"StageKeys,omitempty"`
}

// NewApiGatewayApiKey constructor creates a new ApiGatewayApiKey
func NewApiGatewayApiKey(properties ApiGatewayApiKeyProperties, deps ...interface{}) ApiGatewayApiKey {
	return ApiGatewayApiKey{
		Type:       "AWS::ApiGateway::ApiKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayApiKey parses ApiGatewayApiKey
func ParseApiGatewayApiKey(
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

	// Resources
	var resource ApiGatewayApiKey
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

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-ApiGatewayApiKey-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayApiKey validator
func (resource ApiGatewayApiKey) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayApiKeyProperties validator
func (resource ApiGatewayApiKeyProperties) Validate() []error {
	errors := []error{}
	return errors
}
