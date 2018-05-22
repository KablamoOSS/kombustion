package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRule struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRuleProperties struct {
	MetricName interface{} `yaml:"MetricName"`
	Name interface{} `yaml:"Name"`
	Predicates interface{} `yaml:"Predicates,omitempty"`
}

func NewWAFRule(properties WAFRuleProperties, deps ...interface{}) WAFRule {
	return WAFRule{
		Type:       "AWS::WAF::Rule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRule(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRule - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRuleProperties) Validate() []error {
	errs := []error{}
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
