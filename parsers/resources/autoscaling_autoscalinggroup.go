package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AutoScalingAutoScalingGroup struct {
	Type       string                      `yaml:"Type"`
	Properties AutoScalingAutoScalingGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AutoScalingAutoScalingGroupProperties struct {
	AutoScalingGroupName interface{} `yaml:"AutoScalingGroupName,omitempty"`
	Cooldown interface{} `yaml:"Cooldown,omitempty"`
	DesiredCapacity interface{} `yaml:"DesiredCapacity,omitempty"`
	HealthCheckGracePeriod interface{} `yaml:"HealthCheckGracePeriod,omitempty"`
	HealthCheckType interface{} `yaml:"HealthCheckType,omitempty"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
	LaunchConfigurationName interface{} `yaml:"LaunchConfigurationName,omitempty"`
	MaxSize interface{} `yaml:"MaxSize"`
	MinSize interface{} `yaml:"MinSize"`
	PlacementGroup interface{} `yaml:"PlacementGroup,omitempty"`
	AvailabilityZones interface{} `yaml:"AvailabilityZones,omitempty"`
	LifecycleHookSpecificationList interface{} `yaml:"LifecycleHookSpecificationList,omitempty"`
	LoadBalancerNames interface{} `yaml:"LoadBalancerNames,omitempty"`
	MetricsCollection interface{} `yaml:"MetricsCollection,omitempty"`
	NotificationConfigurations interface{} `yaml:"NotificationConfigurations,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	TargetGroupARNs interface{} `yaml:"TargetGroupARNs,omitempty"`
	TerminationPolicies interface{} `yaml:"TerminationPolicies,omitempty"`
	VPCZoneIdentifier interface{} `yaml:"VPCZoneIdentifier,omitempty"`
}

func NewAutoScalingAutoScalingGroup(properties AutoScalingAutoScalingGroupProperties, deps ...interface{}) AutoScalingAutoScalingGroup {
	return AutoScalingAutoScalingGroup{
		Type:       "AWS::AutoScaling::AutoScalingGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAutoScalingAutoScalingGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource AutoScalingAutoScalingGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingAutoScalingGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AutoScalingAutoScalingGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AutoScalingAutoScalingGroupProperties) Validate() []error {
	errs := []error{}
	if resource.MaxSize == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxSize'"))
	}
	if resource.MinSize == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinSize'"))
	}
	return errs
}
