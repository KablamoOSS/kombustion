package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type SESReceiptRule struct {
	Type       string                      `yaml:"Type"`
	Properties SESReceiptRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SESReceiptRuleProperties struct {
	After interface{} `yaml:"After,omitempty"`
	RuleSetName interface{} `yaml:"RuleSetName"`
	Rule *properties.ReceiptRule_Rule `yaml:"Rule"`
}

func NewSESReceiptRule(properties SESReceiptRuleProperties, deps ...interface{}) SESReceiptRule {
	return SESReceiptRule{
		Type:       "AWS::SES::ReceiptRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESReceiptRule(name string, data string) (cf types.ValueMap, err error) {
	var resource SESReceiptRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESReceiptRule - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SESReceiptRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESReceiptRuleProperties) Validate() []error {
	errs := []error{}
	if resource.RuleSetName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RuleSetName'"))
	}
	if resource.Rule == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rule'"))
	} else {
		errs = append(errs, resource.Rule.Validate()...)
	}
	return errs
}
