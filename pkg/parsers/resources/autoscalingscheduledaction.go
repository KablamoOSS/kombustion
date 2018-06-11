package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// AutoScalingScheduledAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
type AutoScalingScheduledAction struct {
	Type       string                               `yaml:"Type"`
	Properties AutoScalingScheduledActionProperties `yaml:"Properties"`
	Condition  interface{}                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                          `yaml:"DependsOn,omitempty"`
}

// AutoScalingScheduledAction Properties
type AutoScalingScheduledActionProperties struct {
	AutoScalingGroupName interface{} `yaml:"AutoScalingGroupName"`
	DesiredCapacity      interface{} `yaml:"DesiredCapacity,omitempty"`
	EndTime              interface{} `yaml:"EndTime,omitempty"`
	MaxSize              interface{} `yaml:"MaxSize,omitempty"`
	MinSize              interface{} `yaml:"MinSize,omitempty"`
	Recurrence           interface{} `yaml:"Recurrence,omitempty"`
	StartTime            interface{} `yaml:"StartTime,omitempty"`
}

// NewAutoScalingScheduledAction constructor creates a new AutoScalingScheduledAction
func NewAutoScalingScheduledAction(properties AutoScalingScheduledActionProperties, deps ...interface{}) AutoScalingScheduledAction {
	return AutoScalingScheduledAction{
		Type:       "AWS::AutoScaling::ScheduledAction",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAutoScalingScheduledAction parses AutoScalingScheduledAction
func ParseAutoScalingScheduledAction(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
	return
}

// ParseAutoScalingScheduledAction validator
func (resource AutoScalingScheduledAction) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAutoScalingScheduledActionProperties validator
func (resource AutoScalingScheduledActionProperties) Validate() []error {
	errs := []error{}
	if resource.AutoScalingGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoScalingGroupName'"))
	}
	return errs
}