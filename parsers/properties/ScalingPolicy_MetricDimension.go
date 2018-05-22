package properties

	import "fmt"

type ScalingPolicy_MetricDimension struct {
	
	
	Name interface{} `yaml:"Name"`
	Value interface{} `yaml:"Value"`
}

func (resource ScalingPolicy_MetricDimension) Validate() []error {
	errs := []error{}
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
