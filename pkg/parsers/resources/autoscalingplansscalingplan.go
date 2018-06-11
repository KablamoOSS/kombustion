package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// AutoScalingPlansScalingPlan Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html
type AutoScalingPlansScalingPlan struct {
	Type       string                                `yaml:"Type"`
	Properties AutoScalingPlansScalingPlanProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// AutoScalingPlansScalingPlan Properties
type AutoScalingPlansScalingPlanProperties struct {
	ScalingInstructions interface{}                              `yaml:"ScalingInstructions"`
	ApplicationSource   *properties.ScalingPlanApplicationSource `yaml:"ApplicationSource"`
}

// NewAutoScalingPlansScalingPlan constructor creates a new AutoScalingPlansScalingPlan
func NewAutoScalingPlansScalingPlan(properties AutoScalingPlansScalingPlanProperties, deps ...interface{}) AutoScalingPlansScalingPlan {
	return AutoScalingPlansScalingPlan{
		Type:       "AWS::AutoScalingPlans::ScalingPlan",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAutoScalingPlansScalingPlan parses AutoScalingPlansScalingPlan
func ParseAutoScalingPlansScalingPlan(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource AutoScalingPlansScalingPlan
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingPlansScalingPlan - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource AutoScalingPlansScalingPlan) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AutoScalingPlansScalingPlanProperties) Validate() []error {
	errs := []error{}
	if resource.ScalingInstructions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalingInstructions'"))
	}
	if resource.ApplicationSource == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationSource'"))
	} else {
		errs = append(errs, resource.ApplicationSource.Validate()...)
	}
	return errs
}
