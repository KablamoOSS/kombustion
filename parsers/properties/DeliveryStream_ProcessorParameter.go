package properties

	import "fmt"

type DeliveryStream_ProcessorParameter struct {
	
	
	ParameterName interface{} `yaml:"ParameterName"`
	ParameterValue interface{} `yaml:"ParameterValue"`
}

func (resource DeliveryStream_ProcessorParameter) Validate() []error {
	errs := []error{}
	
	
	if resource.ParameterName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterName'"))
	}
	if resource.ParameterValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterValue'"))
	}
	return errs
}
