package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFXssMatchSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFXssMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFXssMatchSetProperties struct {
	Name interface{} `yaml:"Name"`
	XssMatchTuples interface{} `yaml:"XssMatchTuples"`
}

func NewWAFXssMatchSet(properties WAFXssMatchSetProperties, deps ...interface{}) WAFXssMatchSet {
	return WAFXssMatchSet{
		Type:       "AWS::WAF::XssMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFXssMatchSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFXssMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFXssMatchSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFXssMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFXssMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.XssMatchTuples == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'XssMatchTuples'"))
	}
	return errs
}
