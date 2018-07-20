package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFRegionalRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html
type WAFRegionalRule struct {
	Type       string                    `yaml:"Type"`
	Properties WAFRegionalRuleProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// WAFRegionalRule Properties
type WAFRegionalRuleProperties struct {
	MetricName interface{} `yaml:"MetricName"`
	Name       interface{} `yaml:"Name"`
	Predicates interface{} `yaml:"Predicates,omitempty"`
}

// NewWAFRegionalRule constructor creates a new WAFRegionalRule
func NewWAFRegionalRule(properties WAFRegionalRuleProperties, deps ...interface{}) WAFRegionalRule {
	return WAFRegionalRule{
		Type:       "AWS::WAFRegional::Rule",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalRule parses WAFRegionalRule
func ParseWAFRegionalRule(
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
	var resource WAFRegionalRule
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

// ParseWAFRegionalRule validator
func (resource WAFRegionalRule) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalRuleProperties validator
func (resource WAFRegionalRuleProperties) Validate() []error {
	errors := []error{}
	if resource.MetricName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
