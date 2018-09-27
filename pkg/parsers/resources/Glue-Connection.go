package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// GlueConnection Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html
type GlueConnection struct {
	Type       string                   `yaml:"Type"`
	Properties GlueConnectionProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// GlueConnection Properties
type GlueConnectionProperties struct {
	CatalogId       interface{}                           `yaml:"CatalogId"`
	ConnectionInput *properties.ConnectionConnectionInput `yaml:"ConnectionInput"`
}

// NewGlueConnection constructor creates a new GlueConnection
func NewGlueConnection(properties GlueConnectionProperties, deps ...interface{}) GlueConnection {
	return GlueConnection{
		Type:       "AWS::Glue::Connection",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueConnection parses GlueConnection
func ParseGlueConnection(
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
	var resource GlueConnection
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
					"Fn::Sub": "${AWS::StackName}-GlueConnection-" + name,
				},
			},
		},
	}

	return
}

// ParseGlueConnection validator
func (resource GlueConnection) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueConnectionProperties validator
func (resource GlueConnectionProperties) Validate() []error {
	errors := []error{}
	if resource.CatalogId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.ConnectionInput == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ConnectionInput'"))
	} else {
		errors = append(errors, resource.ConnectionInput.Validate()...)
	}
	return errors
}
