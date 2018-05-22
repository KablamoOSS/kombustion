package properties

	import "fmt"

type Pipeline_ActionTypeId struct {
	
	
	
	
	Category interface{} `yaml:"Category"`
	Owner interface{} `yaml:"Owner"`
	Provider interface{} `yaml:"Provider"`
	Version interface{} `yaml:"Version"`
}

func (resource Pipeline_ActionTypeId) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Category == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Category'"))
	}
	if resource.Owner == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Owner'"))
	}
	if resource.Provider == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Provider'"))
	}
	if resource.Version == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Version'"))
	}
	return errs
}
