package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AutoScalingScheduledAction struct {
	Type       string                      `yaml:"Type"`
	Properties AutoScalingScheduledActionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AutoScalingScheduledActionProperties struct {
	AutoScalingGroupName interface{} `yaml:"AutoScalingGroupName"`
	DesiredCapacity interface{} `yaml:"DesiredCapacity,omitempty"`
	EndTime interface{} `yaml:"EndTime,omitempty"`
	MaxSize interface{} `yaml:"MaxSize,omitempty"`
	MinSize interface{} `yaml:"MinSize,omitempty"`
	Recurrence interface{} `yaml:"Recurrence,omitempty"`
	StartTime interface{} `yaml:"StartTime,omitempty"`
}

func NewAutoScalingScheduledAction(properties AutoScalingScheduledActionProperties, deps ...interface{}) AutoScalingScheduledAction {
	return AutoScalingScheduledAction{
		Type:       "AWS::AutoScaling::ScheduledAction",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAutoScalingScheduledAction(name string, data string) (cf types.ValueMap, err error) {
	var resource AutoScalingScheduledAction
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingScheduledAction - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AutoScalingScheduledAction) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AutoScalingScheduledActionProperties) Validate() []error {
	errs := []error{}
	if resource.AutoScalingGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoScalingGroupName'"))
	}
	return errs
}
