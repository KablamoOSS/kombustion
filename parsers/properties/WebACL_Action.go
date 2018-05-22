package properties

	import "fmt"

type WebACL_Action struct {
	
	Type interface{} `yaml:"Type"`
}

func (resource WebACL_Action) Validate() []error {
	errs := []error{}
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
