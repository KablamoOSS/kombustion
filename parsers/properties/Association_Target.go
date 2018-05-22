package properties

	import "fmt"

type Association_Target struct {
	
	
	Key interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values"`
}

func (resource Association_Target) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Values == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Values'"))
	}
	return errs
}
