package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRegionalByteMatchSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalByteMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalByteMatchSetProperties struct {
	Name interface{} `yaml:"Name"`
	ByteMatchTuples interface{} `yaml:"ByteMatchTuples,omitempty"`
}

func NewWAFRegionalByteMatchSet(properties WAFRegionalByteMatchSetProperties, deps ...interface{}) WAFRegionalByteMatchSet {
	return WAFRegionalByteMatchSet{
		Type:       "AWS::WAFRegional::ByteMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalByteMatchSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRegionalByteMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalByteMatchSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalByteMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalByteMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
