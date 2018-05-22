package properties

	import "fmt"

type DeliveryStream_BufferingHints struct {
	
	
	IntervalInSeconds interface{} `yaml:"IntervalInSeconds"`
	SizeInMBs interface{} `yaml:"SizeInMBs"`
}

func (resource DeliveryStream_BufferingHints) Validate() []error {
	errs := []error{}
	
	
	if resource.IntervalInSeconds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IntervalInSeconds'"))
	}
	if resource.SizeInMBs == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SizeInMBs'"))
	}
	return errs
}
