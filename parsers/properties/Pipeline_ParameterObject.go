package properties

	import "fmt"

type Pipeline_ParameterObject struct {
	
	
	Id interface{} `yaml:"Id"`
	Attributes interface{} `yaml:"Attributes"`
}

func (resource Pipeline_ParameterObject) Validate() []error {
	errs := []error{}
	
	
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Attributes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Attributes'"))
	}
	return errs
}
