package properties

	import "fmt"

type Rule_RunCommandTarget struct {
	
	
	Key interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values"`
}

func (resource Rule_RunCommandTarget) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Values == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Values'"))
	}
	return errs
}
