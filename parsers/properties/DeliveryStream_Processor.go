package properties

	import "fmt"

type DeliveryStream_Processor struct {
	
	
	Type interface{} `yaml:"Type"`
	Parameters interface{} `yaml:"Parameters"`
}

func (resource DeliveryStream_Processor) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Parameters == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Parameters'"))
	}
	return errs
}
