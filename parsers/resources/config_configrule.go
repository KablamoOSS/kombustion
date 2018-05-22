package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ConfigConfigRule struct {
	Type       string                      `yaml:"Type"`
	Properties ConfigConfigRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ConfigConfigRuleProperties struct {
	ConfigRuleName interface{} `yaml:"ConfigRuleName,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	InputParameters interface{} `yaml:"InputParameters,omitempty"`
	MaximumExecutionFrequency interface{} `yaml:"MaximumExecutionFrequency,omitempty"`
	Source *properties.ConfigRule_Source `yaml:"Source"`
	Scope *properties.ConfigRule_Scope `yaml:"Scope,omitempty"`
}

func NewConfigConfigRule(properties ConfigConfigRuleProperties, deps ...interface{}) ConfigConfigRule {
	return ConfigConfigRule{
		Type:       "AWS::Config::ConfigRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseConfigConfigRule(name string, data string) (cf types.ValueMap, err error) {
	var resource ConfigConfigRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ConfigConfigRule - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ConfigConfigRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ConfigConfigRuleProperties) Validate() []error {
	errs := []error{}
	if resource.Source == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Source'"))
	} else {
		errs = append(errs, resource.Source.Validate()...)
	}
	return errs
}
