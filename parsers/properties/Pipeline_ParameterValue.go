package properties

	import "fmt"

type Pipeline_ParameterValue struct {
	
	
	Id interface{} `yaml:"Id"`
	StringValue interface{} `yaml:"StringValue"`
}

func (resource Pipeline_ParameterValue) Validate() []error {
	errs := []error{}
	
	
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.StringValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StringValue'"))
	}
	return errs
}
