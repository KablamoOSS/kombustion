package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFIPSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFIPSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFIPSetProperties struct {
	Name interface{} `yaml:"Name"`
	IPSetDescriptors interface{} `yaml:"IPSetDescriptors,omitempty"`
}

func NewWAFIPSet(properties WAFIPSetProperties, deps ...interface{}) WAFIPSet {
	return WAFIPSet{
		Type:       "AWS::WAF::IPSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFIPSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFIPSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFIPSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFIPSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFIPSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
