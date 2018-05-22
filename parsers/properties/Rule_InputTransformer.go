package properties

	import "fmt"

type Rule_InputTransformer struct {
	
	
	InputTemplate interface{} `yaml:"InputTemplate"`
	InputPathsMap interface{} `yaml:"InputPathsMap,omitempty"`
}

func (resource Rule_InputTransformer) Validate() []error {
	errs := []error{}
	
	
	if resource.InputTemplate == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InputTemplate'"))
	}
	return errs
}
