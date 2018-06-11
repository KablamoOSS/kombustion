package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
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
func ParseWAFRegionalRule(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFRegionalRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalRule - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseWAFRegionalRule validator
func (resource WAFRegionalRule) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalRuleProperties validator
func (resource WAFRegionalRuleProperties) Validate() []error {
	errs := []error{}
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}