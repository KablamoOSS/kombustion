package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AutoScalingLifecycleHook struct {
	Type       string                      `yaml:"Type"`
	Properties AutoScalingLifecycleHookProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AutoScalingLifecycleHookProperties struct {
	AutoScalingGroupName interface{} `yaml:"AutoScalingGroupName"`
	DefaultResult interface{} `yaml:"DefaultResult,omitempty"`
	HeartbeatTimeout interface{} `yaml:"HeartbeatTimeout,omitempty"`
	LifecycleHookName interface{} `yaml:"LifecycleHookName,omitempty"`
	LifecycleTransition interface{} `yaml:"LifecycleTransition"`
	NotificationMetadata interface{} `yaml:"NotificationMetadata,omitempty"`
	NotificationTargetARN interface{} `yaml:"NotificationTargetARN,omitempty"`
	RoleARN interface{} `yaml:"RoleARN,omitempty"`
}

func NewAutoScalingLifecycleHook(properties AutoScalingLifecycleHookProperties, deps ...interface{}) AutoScalingLifecycleHook {
	return AutoScalingLifecycleHook{
		Type:       "AWS::AutoScaling::LifecycleHook",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAutoScalingLifecycleHook(name string, data string) (cf types.ValueMap, err error) {
	var resource AutoScalingLifecycleHook
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingLifecycleHook - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AutoScalingLifecycleHook) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AutoScalingLifecycleHookProperties) Validate() []error {
	errs := []error{}
	if resource.AutoScalingGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoScalingGroupName'"))
	}
	if resource.LifecycleTransition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LifecycleTransition'"))
	}
	return errs
}
