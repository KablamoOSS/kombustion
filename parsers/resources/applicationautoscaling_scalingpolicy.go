package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ApplicationAutoScalingScalingPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties ApplicationAutoScalingScalingPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApplicationAutoScalingScalingPolicyProperties struct {
	PolicyName interface{} `yaml:"PolicyName"`
	PolicyType interface{} `yaml:"PolicyType"`
	ResourceId interface{} `yaml:"ResourceId,omitempty"`
	ScalableDimension interface{} `yaml:"ScalableDimension,omitempty"`
	ScalingTargetId interface{} `yaml:"ScalingTargetId,omitempty"`
	ServiceNamespace interface{} `yaml:"ServiceNamespace,omitempty"`
	TargetTrackingScalingPolicyConfiguration *properties.ScalingPolicy_TargetTrackingScalingPolicyConfiguration `yaml:"TargetTrackingScalingPolicyConfiguration,omitempty"`
	StepScalingPolicyConfiguration *properties.ScalingPolicy_StepScalingPolicyConfiguration `yaml:"StepScalingPolicyConfiguration,omitempty"`
}

func NewApplicationAutoScalingScalingPolicy(properties ApplicationAutoScalingScalingPolicyProperties, deps ...interface{}) ApplicationAutoScalingScalingPolicy {
	return ApplicationAutoScalingScalingPolicy{
		Type:       "AWS::ApplicationAutoScaling::ScalingPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApplicationAutoScalingScalingPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource ApplicationAutoScalingScalingPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApplicationAutoScalingScalingPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApplicationAutoScalingScalingPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApplicationAutoScalingScalingPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	if resource.PolicyType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyType'"))
	}
	return errs
}
