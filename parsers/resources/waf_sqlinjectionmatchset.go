package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFSqlInjectionMatchSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFSqlInjectionMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFSqlInjectionMatchSetProperties struct {
	Name interface{} `yaml:"Name"`
	SqlInjectionMatchTuples interface{} `yaml:"SqlInjectionMatchTuples,omitempty"`
}

func NewWAFSqlInjectionMatchSet(properties WAFSqlInjectionMatchSetProperties, deps ...interface{}) WAFSqlInjectionMatchSet {
	return WAFSqlInjectionMatchSet{
		Type:       "AWS::WAF::SqlInjectionMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFSqlInjectionMatchSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFSqlInjectionMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFSqlInjectionMatchSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFSqlInjectionMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFSqlInjectionMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
