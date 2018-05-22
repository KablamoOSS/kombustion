package properties

	import "fmt"

type Pipeline_Field struct {
	
	
	
	Key interface{} `yaml:"Key"`
	RefValue interface{} `yaml:"RefValue,omitempty"`
	StringValue interface{} `yaml:"StringValue,omitempty"`
}

func (resource Pipeline_Field) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}
