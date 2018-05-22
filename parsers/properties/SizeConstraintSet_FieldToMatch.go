package properties

	import "fmt"

type SizeConstraintSet_FieldToMatch struct {
	
	
	Data interface{} `yaml:"Data,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource SizeConstraintSet_FieldToMatch) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
