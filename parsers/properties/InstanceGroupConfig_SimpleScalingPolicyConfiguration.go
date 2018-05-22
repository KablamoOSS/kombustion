package properties

	import "fmt"

type InstanceGroupConfig_SimpleScalingPolicyConfiguration struct {
	
	
	
	AdjustmentType interface{} `yaml:"AdjustmentType,omitempty"`
	CoolDown interface{} `yaml:"CoolDown,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment"`
}

func (resource InstanceGroupConfig_SimpleScalingPolicyConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.ScalingAdjustment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalingAdjustment'"))
	}
	return errs
}
