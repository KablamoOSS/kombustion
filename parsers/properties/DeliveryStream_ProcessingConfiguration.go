package properties

	import "fmt"

type DeliveryStream_ProcessingConfiguration struct {
	
	
	Enabled interface{} `yaml:"Enabled"`
	Processors interface{} `yaml:"Processors"`
}

func (resource DeliveryStream_ProcessingConfiguration) Validate() []error {
	errs := []error{}
	
	
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	if resource.Processors == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Processors'"))
	}
	return errs
}
