package properties


type Layer_AutoScalingThresholds struct {
	
	
	
	
	
	
	CpuThreshold interface{} `yaml:"CpuThreshold,omitempty"`
	IgnoreMetricsTime interface{} `yaml:"IgnoreMetricsTime,omitempty"`
	InstanceCount interface{} `yaml:"InstanceCount,omitempty"`
	LoadThreshold interface{} `yaml:"LoadThreshold,omitempty"`
	MemoryThreshold interface{} `yaml:"MemoryThreshold,omitempty"`
	ThresholdsWaitTime interface{} `yaml:"ThresholdsWaitTime,omitempty"`
}

func (resource Layer_AutoScalingThresholds) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
