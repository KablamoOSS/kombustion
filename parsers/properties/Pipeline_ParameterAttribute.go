package properties

	import "fmt"

type Pipeline_ParameterAttribute struct {
	
	
	Key interface{} `yaml:"Key"`
	StringValue interface{} `yaml:"StringValue"`
}

func (resource Pipeline_ParameterAttribute) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.StringValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StringValue'"))
	}
	return errs
}
