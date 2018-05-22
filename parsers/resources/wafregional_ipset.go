package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRegionalIPSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalIPSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalIPSetProperties struct {
	Name interface{} `yaml:"Name"`
	IPSetDescriptors interface{} `yaml:"IPSetDescriptors,omitempty"`
}

func NewWAFRegionalIPSet(properties WAFRegionalIPSetProperties, deps ...interface{}) WAFRegionalIPSet {
	return WAFRegionalIPSet{
		Type:       "AWS::WAFRegional::IPSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalIPSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRegionalIPSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalIPSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalIPSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalIPSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
