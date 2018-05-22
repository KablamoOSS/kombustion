package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type WAFRegionalSizeConstraintSet struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalSizeConstraintSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type WAFRegionalSizeConstraintSetProperties struct {
	Name interface{} `yaml:"Name"`
	SizeConstraints interface{} `yaml:"SizeConstraints,omitempty"`
}

func NewWAFRegionalSizeConstraintSet(properties WAFRegionalSizeConstraintSetProperties, deps ...interface{}) WAFRegionalSizeConstraintSet {
	return WAFRegionalSizeConstraintSet{
		Type:       "AWS::WAFRegional::SizeConstraintSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalSizeConstraintSet(name string, data string) (cf types.ValueMap, err error) {
	var resource WAFRegionalSizeConstraintSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalSizeConstraintSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource WAFRegionalSizeConstraintSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalSizeConstraintSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
