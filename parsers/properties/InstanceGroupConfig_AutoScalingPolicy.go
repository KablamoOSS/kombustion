package properties

	import "fmt"

type InstanceGroupConfig_AutoScalingPolicy struct {
	
	
	Constraints *InstanceGroupConfig_ScalingConstraints `yaml:"Constraints"`
	Rules interface{} `yaml:"Rules"`
}

func (resource InstanceGroupConfig_AutoScalingPolicy) Validate() []error {
	errs := []error{}
	
	
	if resource.Constraints == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Constraints'"))
	} else {
		errs = append(errs, resource.Constraints.Validate()...)
	}
	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
