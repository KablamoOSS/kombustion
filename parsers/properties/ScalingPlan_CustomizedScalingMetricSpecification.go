package properties

	import "fmt"

type ScalingPlan_CustomizedScalingMetricSpecification struct {
	
	
	
	
	
	MetricName interface{} `yaml:"MetricName"`
	Namespace interface{} `yaml:"Namespace"`
	Statistic interface{} `yaml:"Statistic"`
	Unit interface{} `yaml:"Unit,omitempty"`
	Dimensions interface{} `yaml:"Dimensions,omitempty"`
}

func (resource ScalingPlan_CustomizedScalingMetricSpecification) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Namespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Namespace'"))
	}
	if resource.Statistic == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Statistic'"))
	}
	return errs
}
