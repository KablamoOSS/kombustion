package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DirectoryServiceSimpleAD Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html
type DirectoryServiceSimpleAD struct {
	Type       string                             `yaml:"Type"`
	Properties DirectoryServiceSimpleADProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// DirectoryServiceSimpleAD Properties
type DirectoryServiceSimpleADProperties struct {
	CreateAlias interface{}                     `yaml:"CreateAlias,omitempty"`
	Description interface{}                     `yaml:"Description,omitempty"`
	EnableSso   interface{}                     `yaml:"EnableSso,omitempty"`
	Name        interface{}                     `yaml:"Name"`
	Password    interface{}                     `yaml:"Password"`
	ShortName   interface{}                     `yaml:"ShortName,omitempty"`
	Size        interface{}                     `yaml:"Size"`
	VpcSettings *properties.SimpleADVpcSettings `yaml:"VpcSettings"`
}

// NewDirectoryServiceSimpleAD constructor creates a new DirectoryServiceSimpleAD
func NewDirectoryServiceSimpleAD(properties DirectoryServiceSimpleADProperties, deps ...interface{}) DirectoryServiceSimpleAD {
	return DirectoryServiceSimpleAD{
		Type:       "AWS::DirectoryService::SimpleAD",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDirectoryServiceSimpleAD parses DirectoryServiceSimpleAD
func ParseDirectoryServiceSimpleAD(
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
	var resource DirectoryServiceSimpleAD
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
					"Fn::Sub": "${AWS::StackName}-DirectoryServiceSimpleAD-" + name,
				},
			},
		},
	}

	return
}

// ParseDirectoryServiceSimpleAD validator
func (resource DirectoryServiceSimpleAD) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDirectoryServiceSimpleADProperties validator
func (resource DirectoryServiceSimpleADProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Password == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Password'"))
	}
	if resource.Size == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Size'"))
	}
	if resource.VpcSettings == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'VpcSettings'"))
	} else {
		errors = append(errors, resource.VpcSettings.Validate()...)
	}
	return errors
}
