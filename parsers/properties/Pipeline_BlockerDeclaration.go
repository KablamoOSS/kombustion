package properties

	import "fmt"

type Pipeline_BlockerDeclaration struct {
	
	
	Name interface{} `yaml:"Name"`
	Type interface{} `yaml:"Type"`
}

func (resource Pipeline_BlockerDeclaration) Validate() []error {
	errs := []error{}
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
