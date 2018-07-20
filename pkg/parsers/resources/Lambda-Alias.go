package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// LambdaAlias Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html
type LambdaAlias struct {
	Type       string                `yaml:"Type"`
	Properties LambdaAliasProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// LambdaAlias Properties
type LambdaAliasProperties struct {
	Description     interface{}                                `yaml:"Description,omitempty"`
	FunctionName    interface{}                                `yaml:"FunctionName"`
	FunctionVersion interface{}                                `yaml:"FunctionVersion"`
	Name            interface{}                                `yaml:"Name"`
	RoutingConfig   *properties.AliasAliasRoutingConfiguration `yaml:"RoutingConfig,omitempty"`
}

// NewLambdaAlias constructor creates a new LambdaAlias
func NewLambdaAlias(properties LambdaAliasProperties, deps ...interface{}) LambdaAlias {
	return LambdaAlias{
		Type:       "AWS::Lambda::Alias",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLambdaAlias parses LambdaAlias
func ParseLambdaAlias(
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
	var resource LambdaAlias
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

// ParseLambdaAlias validator
func (resource LambdaAlias) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLambdaAliasProperties validator
func (resource LambdaAliasProperties) Validate() []error {
	errors := []error{}
	if resource.FunctionName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	if resource.FunctionVersion == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'FunctionVersion'"))
	}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
