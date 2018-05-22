package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type EventsRule struct {
	Type       string                      `yaml:"Type"`
	Properties EventsRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EventsRuleProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	EventPattern interface{} `yaml:"EventPattern,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	RoleArn interface{} `yaml:"RoleArn,omitempty"`
	ScheduleExpression interface{} `yaml:"ScheduleExpression,omitempty"`
	State interface{} `yaml:"State,omitempty"`
	Targets interface{} `yaml:"Targets,omitempty"`
}

func NewEventsRule(properties EventsRuleProperties, deps ...interface{}) EventsRule {
	return EventsRule{
		Type:       "AWS::Events::Rule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEventsRule(name string, data string) (cf types.ValueMap, err error) {
	var resource EventsRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EventsRule - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EventsRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EventsRuleProperties) Validate() []error {
	errs := []error{}
	return errs
}
