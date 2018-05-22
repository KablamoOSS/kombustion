package properties

	import "fmt"

type ScalingPolicy_PredefinedMetricSpecification struct {
	
	
	PredefinedMetricType interface{} `yaml:"PredefinedMetricType"`
	ResourceLabel interface{} `yaml:"ResourceLabel,omitempty"`
}

func (resource ScalingPolicy_PredefinedMetricSpecification) Validate() []error {
	errs := []error{}
	
	
	if resource.PredefinedMetricType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PredefinedMetricType'"))
	}
	return errs
}
