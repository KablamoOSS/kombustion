package properties

	import "fmt"

type InstanceGroupConfig_ScalingRule struct {
	
	
	
	
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	Trigger *InstanceGroupConfig_ScalingTrigger `yaml:"Trigger"`
	Action *InstanceGroupConfig_ScalingAction `yaml:"Action"`
}

func (resource InstanceGroupConfig_ScalingRule) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Trigger == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Trigger'"))
	} else {
		errs = append(errs, resource.Trigger.Validate()...)
	}
	if resource.Action == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Action'"))
	} else {
		errs = append(errs, resource.Action.Validate()...)
	}
	return errs
}
