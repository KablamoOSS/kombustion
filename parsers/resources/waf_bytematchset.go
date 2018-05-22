package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFByteMatchSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFByteMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFByteMatchSetProperties struct {
	Name interface{} `yaml:"Name"`
	ByteMatchTuples interface{} `yaml:"ByteMatchTuples,omitempty"`
}

func NewWAFByteMatchSet(properties WAFByteMatchSetProperties, deps ...interface{}) WAFByteMatchSet {
	return WAFByteMatchSet{
		Type:       "AWS::WAF::ByteMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFByteMatchSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFByteMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFByteMatchSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFByteMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFByteMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
