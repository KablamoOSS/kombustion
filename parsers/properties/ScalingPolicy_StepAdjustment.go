package properties

	import "fmt"

type ScalingPolicy_StepAdjustment struct {
	
	
	
	MetricIntervalLowerBound interface{} `yaml:"MetricIntervalLowerBound,omitempty"`
	MetricIntervalUpperBound interface{} `yaml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment"`
}

func (resource ScalingPolicy_StepAdjustment) Validate() []error {
	errs := []error{}
	
	
	
	if resource.ScalingAdjustment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalingAdjustment'"))
	}
	return errs
}
