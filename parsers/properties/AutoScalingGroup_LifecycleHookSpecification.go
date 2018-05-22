package properties

	import "fmt"

type AutoScalingGroup_LifecycleHookSpecification struct {
	
	
	
	
	
	
	
	DefaultResult interface{} `yaml:"DefaultResult,omitempty"`
	HeartbeatTimeout interface{} `yaml:"HeartbeatTimeout,omitempty"`
	LifecycleHookName interface{} `yaml:"LifecycleHookName"`
	LifecycleTransition interface{} `yaml:"LifecycleTransition"`
	NotificationMetadata interface{} `yaml:"NotificationMetadata,omitempty"`
	NotificationTargetARN interface{} `yaml:"NotificationTargetARN,omitempty"`
	RoleARN interface{} `yaml:"RoleARN,omitempty"`
}

func (resource AutoScalingGroup_LifecycleHookSpecification) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	if resource.LifecycleHookName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LifecycleHookName'"))
	}
	if resource.LifecycleTransition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LifecycleTransition'"))
	}
	return errs
}
