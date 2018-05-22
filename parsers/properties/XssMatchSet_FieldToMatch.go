package properties

	import "fmt"

type XssMatchSet_FieldToMatch struct {
	
	
	Data interface{} `yaml:"Data,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource XssMatchSet_FieldToMatch) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
