package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueTrigger struct {
	Type       string                      `yaml:"Type"`
	Properties GlueTriggerProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueTriggerProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Schedule interface{} `yaml:"Schedule,omitempty"`
	Type interface{} `yaml:"Type"`
	Predicate *properties.Trigger_Predicate `yaml:"Predicate,omitempty"`
	Actions interface{} `yaml:"Actions"`
}

func NewGlueTrigger(properties GlueTriggerProperties, deps ...interface{}) GlueTrigger {
	return GlueTrigger{
		Type:       "AWS::Glue::Trigger",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueTrigger(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueTrigger
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueTrigger - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueTrigger) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueTriggerProperties) Validate() []error {
	errs := []error{}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Actions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Actions'"))
	}
	return errs
}
