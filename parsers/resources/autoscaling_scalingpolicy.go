package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type AutoScalingScalingPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties AutoScalingScalingPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AutoScalingScalingPolicyProperties struct {
	AdjustmentType interface{} `yaml:"AdjustmentType,omitempty"`
	AutoScalingGroupName interface{} `yaml:"AutoScalingGroupName"`
	Cooldown interface{} `yaml:"Cooldown,omitempty"`
	EstimatedInstanceWarmup interface{} `yaml:"EstimatedInstanceWarmup,omitempty"`
	MetricAggregationType interface{} `yaml:"MetricAggregationType,omitempty"`
	MinAdjustmentMagnitude interface{} `yaml:"MinAdjustmentMagnitude,omitempty"`
	PolicyType interface{} `yaml:"PolicyType,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment,omitempty"`
	TargetTrackingConfiguration *properties.ScalingPolicy_TargetTrackingConfiguration `yaml:"TargetTrackingConfiguration,omitempty"`
	StepAdjustments interface{} `yaml:"StepAdjustments,omitempty"`
}

func NewAutoScalingScalingPolicy(properties AutoScalingScalingPolicyProperties, deps ...interface{}) AutoScalingScalingPolicy {
	return AutoScalingScalingPolicy{
		Type:       "AWS::AutoScaling::ScalingPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAutoScalingScalingPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource AutoScalingScalingPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingScalingPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AutoScalingScalingPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AutoScalingScalingPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.AutoScalingGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoScalingGroupName'"))
	}
	return errs
}
