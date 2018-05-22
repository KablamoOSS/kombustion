package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type SESReceiptFilter struct {
	Type       string                      `yaml:"Type"`
	Properties SESReceiptFilterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SESReceiptFilterProperties struct {
	Filter *properties.ReceiptFilter_Filter `yaml:"Filter"`
}

func NewSESReceiptFilter(properties SESReceiptFilterProperties, deps ...interface{}) SESReceiptFilter {
	return SESReceiptFilter{
		Type:       "AWS::SES::ReceiptFilter",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESReceiptFilter(name string, data string) (cf types.ValueMap, err error) {
	var resource SESReceiptFilter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESReceiptFilter - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SESReceiptFilter) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESReceiptFilterProperties) Validate() []error {
	errs := []error{}
	if resource.Filter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Filter'"))
	} else {
		errs = append(errs, resource.Filter.Validate()...)
	}
	return errs
}
