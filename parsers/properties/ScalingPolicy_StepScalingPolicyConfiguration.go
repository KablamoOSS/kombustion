package properties


type ScalingPolicy_StepScalingPolicyConfiguration struct {
	
	
	
	
	
	AdjustmentType interface{} `yaml:"AdjustmentType,omitempty"`
	Cooldown interface{} `yaml:"Cooldown,omitempty"`
	MetricAggregationType interface{} `yaml:"MetricAggregationType,omitempty"`
	MinAdjustmentMagnitude interface{} `yaml:"MinAdjustmentMagnitude,omitempty"`
	StepAdjustments interface{} `yaml:"StepAdjustments,omitempty"`
}

func (resource ScalingPolicy_StepScalingPolicyConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
