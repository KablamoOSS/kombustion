package properties

	import "fmt"

type Association_ParameterValues struct {
	
	ParameterValues interface{} `yaml:"ParameterValues"`
}

func (resource Association_ParameterValues) Validate() []error {
	errs := []error{}
	
	if resource.ParameterValues == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterValues'"))
	}
	return errs
}
