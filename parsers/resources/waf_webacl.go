package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type WAFWebACL struct {
	Type       string                      `yaml:"Type"`
	Properties WAFWebACLProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFWebACLProperties struct {
	MetricName interface{} `yaml:"MetricName"`
	Name interface{} `yaml:"Name"`
	DefaultAction *properties.WebACL_WafAction `yaml:"DefaultAction"`
	Rules interface{} `yaml:"Rules,omitempty"`
}

func NewWAFWebACL(properties WAFWebACLProperties, deps ...interface{}) WAFWebACL {
	return WAFWebACL{
		Type:       "AWS::WAF::WebACL",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFWebACL(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFWebACL
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFWebACL - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFWebACL) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFWebACLProperties) Validate() []error {
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
