package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRegionalRule struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalRuleProperties struct {
	MetricName interface{} `yaml:"MetricName"`
	Name interface{} `yaml:"Name"`
	Predicates interface{} `yaml:"Predicates,omitempty"`
}

func NewWAFRegionalRule(properties WAFRegionalRuleProperties, deps ...interface{}) WAFRegionalRule {
	return WAFRegionalRule{
		Type:       "AWS::WAFRegional::Rule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalRule(name string, data string) (cf types.ValueMap, err error) {
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
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalRule) Validate() []error {
	return resource.Properties.Validate()
}

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
