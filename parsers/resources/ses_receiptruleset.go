package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type SESReceiptRuleSet struct {
	Type       string                      `yaml:"Type"`
	Properties SESReceiptRuleSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SESReceiptRuleSetProperties struct {
	RuleSetName interface{} `yaml:"RuleSetName,omitempty"`
}

func NewSESReceiptRuleSet(properties SESReceiptRuleSetProperties, deps ...interface{}) SESReceiptRuleSet {
	return SESReceiptRuleSet{
		Type:       "AWS::SES::ReceiptRuleSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESReceiptRuleSet(name string, data string) (cf types.ValueMap, err error) {
	var resource SESReceiptRuleSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESReceiptRuleSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SESReceiptRuleSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESReceiptRuleSetProperties) Validate() []error {
	errs := []error{}
	return errs
}
