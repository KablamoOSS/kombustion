package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFIPSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html
type WAFIPSet struct {
	Type       string             `yaml:"Type"`
	Properties WAFIPSetProperties `yaml:"Properties"`
	Condition  interface{}        `yaml:"Condition,omitempty"`
	Metadata   interface{}        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}        `yaml:"DependsOn,omitempty"`
}

// WAFIPSet Properties
type WAFIPSetProperties struct {
	Name             interface{} `yaml:"Name"`
	IPSetDescriptors interface{} `yaml:"IPSetDescriptors,omitempty"`
}

// NewWAFIPSet constructor creates a new WAFIPSet
func NewWAFIPSet(properties WAFIPSetProperties, deps ...interface{}) WAFIPSet {
	return WAFIPSet{
		Type:       "AWS::WAF::IPSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFIPSet parses WAFIPSet
func ParseWAFIPSet(
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
	var resource WAFIPSet
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
					"Fn::Sub": "${AWS::StackName}-WAFIPSet-" + name,
				},
			},
		},
	}

	return
}

// ParseWAFIPSet validator
func (resource WAFIPSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFIPSetProperties validator
func (resource WAFIPSetProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
