package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApplicationAutoScalingScalableTarget struct {
	Type       string                      `yaml:"Type"`
	Properties ApplicationAutoScalingScalableTargetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApplicationAutoScalingScalableTargetProperties struct {
	MaxCapacity interface{} `yaml:"MaxCapacity"`
	MinCapacity interface{} `yaml:"MinCapacity"`
	ResourceId interface{} `yaml:"ResourceId"`
	RoleARN interface{} `yaml:"RoleARN"`
	ScalableDimension interface{} `yaml:"ScalableDimension"`
	ServiceNamespace interface{} `yaml:"ServiceNamespace"`
	ScheduledActions interface{} `yaml:"ScheduledActions,omitempty"`
}

func NewApplicationAutoScalingScalableTarget(properties ApplicationAutoScalingScalableTargetProperties, deps ...interface{}) ApplicationAutoScalingScalableTarget {
	return ApplicationAutoScalingScalableTarget{
		Type:       "AWS::ApplicationAutoScaling::ScalableTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApplicationAutoScalingScalableTarget(name string, data string) (cf types.ValueMap, err error) {
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
	cf = types.ValueMap{name: resource}
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
