package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRegionalSqlInjectionMatchSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalSqlInjectionMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalSqlInjectionMatchSetProperties struct {
	Name interface{} `yaml:"Name"`
	SqlInjectionMatchTuples interface{} `yaml:"SqlInjectionMatchTuples,omitempty"`
}

func NewWAFRegionalSqlInjectionMatchSet(properties WAFRegionalSqlInjectionMatchSetProperties, deps ...interface{}) WAFRegionalSqlInjectionMatchSet {
	return WAFRegionalSqlInjectionMatchSet{
		Type:       "AWS::WAFRegional::SqlInjectionMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalSqlInjectionMatchSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRegionalSqlInjectionMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalSqlInjectionMatchSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalSqlInjectionMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalSqlInjectionMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
