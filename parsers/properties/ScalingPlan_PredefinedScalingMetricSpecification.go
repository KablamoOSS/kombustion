package properties

	import "fmt"

type ScalingPlan_PredefinedScalingMetricSpecification struct {
	
	
	PredefinedScalingMetricType interface{} `yaml:"PredefinedScalingMetricType"`
	ResourceLabel interface{} `yaml:"ResourceLabel,omitempty"`
}

func (resource ScalingPlan_PredefinedScalingMetricSpecification) Validate() []error {
	errs := []error{}
	
	
	if resource.PredefinedScalingMetricType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PredefinedScalingMetricType'"))
	}
	return errs
}
