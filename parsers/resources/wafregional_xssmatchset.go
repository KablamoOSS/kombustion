package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRegionalXssMatchSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalXssMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalXssMatchSetProperties struct {
	Name interface{} `yaml:"Name"`
	XssMatchTuples interface{} `yaml:"XssMatchTuples,omitempty"`
}

func NewWAFRegionalXssMatchSet(properties WAFRegionalXssMatchSetProperties, deps ...interface{}) WAFRegionalXssMatchSet {
	return WAFRegionalXssMatchSet{
		Type:       "AWS::WAFRegional::XssMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalXssMatchSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRegionalXssMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalXssMatchSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalXssMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalXssMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
