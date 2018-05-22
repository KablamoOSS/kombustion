package properties

	import "fmt"

type IPSet_IPSetDescriptor struct {
	
	
	Type interface{} `yaml:"Type"`
	Value interface{} `yaml:"Value"`
}

func (resource IPSet_IPSetDescriptor) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
