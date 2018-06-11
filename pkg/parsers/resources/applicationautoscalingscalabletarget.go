package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApplicationAutoScalingScalableTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html
type ApplicationAutoScalingScalableTarget struct {
	Type       string                                         `yaml:"Type"`
	Properties ApplicationAutoScalingScalableTargetProperties `yaml:"Properties"`
	Condition  interface{}                                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                    `yaml:"DependsOn,omitempty"`
}

// ApplicationAutoScalingScalableTarget Properties
type ApplicationAutoScalingScalableTargetProperties struct {
	MaxCapacity       interface{} `yaml:"MaxCapacity"`
	MinCapacity       interface{} `yaml:"MinCapacity"`
	ResourceId        interface{} `yaml:"ResourceId"`
	RoleARN           interface{} `yaml:"RoleARN"`
	ScalableDimension interface{} `yaml:"ScalableDimension"`
	ServiceNamespace  interface{} `yaml:"ServiceNamespace"`
	ScheduledActions  interface{} `yaml:"ScheduledActions,omitempty"`
}

// NewApplicationAutoScalingScalableTarget constructor creates a new ApplicationAutoScalingScalableTarget
func NewApplicationAutoScalingScalableTarget(properties ApplicationAutoScalingScalableTargetProperties, deps ...interface{}) ApplicationAutoScalingScalableTarget {
	return ApplicationAutoScalingScalableTarget{
		Type:       "AWS::ApplicationAutoScaling::ScalableTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApplicationAutoScalingScalableTarget parses ApplicationAutoScalingScalableTarget
func ParseApplicationAutoScalingScalableTarget(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApplicationAutoScalingScalableTarget
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApplicationAutoScalingScalableTarget - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ApplicationAutoScalingScalableTarget) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApplicationAutoScalingScalableTargetProperties) Validate() []error {
	errs := []error{}
	if resource.MaxCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxCapacity'"))
	}
	if resource.MinCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinCapacity'"))
	}
	if resource.ResourceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceId'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.ScalableDimension == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalableDimension'"))
	}
	if resource.ServiceNamespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceNamespace'"))
	}
	return errs
}
