package properties

	import "fmt"

type Project_SourceAuth struct {
	
	
	Resource interface{} `yaml:"Resource,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource Project_SourceAuth) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
