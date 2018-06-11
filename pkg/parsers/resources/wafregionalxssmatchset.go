package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// WAFRegionalXssMatchSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html
type WAFRegionalXssMatchSet struct {
	Type       string                           `yaml:"Type"`
	Properties WAFRegionalXssMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                      `yaml:"DependsOn,omitempty"`
}

// WAFRegionalXssMatchSet Properties
type WAFRegionalXssMatchSetProperties struct {
	Name           interface{} `yaml:"Name"`
	XssMatchTuples interface{} `yaml:"XssMatchTuples,omitempty"`
}

// NewWAFRegionalXssMatchSet constructor creates a new WAFRegionalXssMatchSet
func NewWAFRegionalXssMatchSet(properties WAFRegionalXssMatchSetProperties, deps ...interface{}) WAFRegionalXssMatchSet {
	return WAFRegionalXssMatchSet{
		Type:       "AWS::WAFRegional::XssMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalXssMatchSet parses WAFRegionalXssMatchSet
func ParseWAFRegionalXssMatchSet(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFRegionalXssMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalXssMatchSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource WAFRegionalXssMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalXssMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
