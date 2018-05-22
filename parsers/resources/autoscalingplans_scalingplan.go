package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type AutoScalingPlansScalingPlan struct {
	Type       string                      `yaml:"Type"`
	Properties AutoScalingPlansScalingPlanProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AutoScalingPlansScalingPlanProperties struct {
	ScalingInstructions interface{} `yaml:"ScalingInstructions"`
	ApplicationSource *properties.ScalingPlan_ApplicationSource `yaml:"ApplicationSource"`
}

func NewAutoScalingPlansScalingPlan(properties AutoScalingPlansScalingPlanProperties, deps ...interface{}) AutoScalingPlansScalingPlan {
	return AutoScalingPlansScalingPlan{
		Type:       "AWS::AutoScalingPlans::ScalingPlan",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAutoScalingPlansScalingPlan(name string, data string) (cf types.ValueMap, err error) {
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
	cf = types.ValueMap{name: resource}
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
