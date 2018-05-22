package properties

	import "fmt"

type Cluster_SimpleScalingPolicyConfiguration struct {
	
	
	
	AdjustmentType interface{} `yaml:"AdjustmentType,omitempty"`
	CoolDown interface{} `yaml:"CoolDown,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment"`
}

func (resource Cluster_SimpleScalingPolicyConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.ScalingAdjustment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalingAdjustment'"))
	}
	return errs
}
