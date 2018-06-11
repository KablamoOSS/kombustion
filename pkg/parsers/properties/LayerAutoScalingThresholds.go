package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LayerAutoScalingThresholds Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html
type LayerAutoScalingThresholds struct {
	CpuThreshold       interface{} `yaml:"CpuThreshold,omitempty"`
	IgnoreMetricsTime  interface{} `yaml:"IgnoreMetricsTime,omitempty"`
	InstanceCount      interface{} `yaml:"InstanceCount,omitempty"`
	LoadThreshold      interface{} `yaml:"LoadThreshold,omitempty"`
	MemoryThreshold    interface{} `yaml:"MemoryThreshold,omitempty"`
	ThresholdsWaitTime interface{} `yaml:"ThresholdsWaitTime,omitempty"`
}

func (resource LayerAutoScalingThresholds) Validate() []error {
	errs := []error{}

	return errs
}
