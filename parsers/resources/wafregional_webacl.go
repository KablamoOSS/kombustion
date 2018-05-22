package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type WAFRegionalWebACL struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalWebACLProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalWebACLProperties struct {
	MetricName interface{} `yaml:"MetricName"`
	Name interface{} `yaml:"Name"`
	Rules interface{} `yaml:"Rules,omitempty"`
	DefaultAction *properties.WebACL_Action `yaml:"DefaultAction"`
}

func NewWAFRegionalWebACL(properties WAFRegionalWebACLProperties, deps ...interface{}) WAFRegionalWebACL {
	return WAFRegionalWebACL{
		Type:       "AWS::WAFRegional::WebACL",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalWebACL(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRegionalWebACL
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalWebACL - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalWebACL) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalWebACLProperties) Validate() []error {
	errs := []error{}
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.DefaultAction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DefaultAction'"))
	} else {
		errs = append(errs, resource.DefaultAction.Validate()...)
	}
	return errs
}
