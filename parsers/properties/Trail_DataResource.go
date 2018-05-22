package properties

	import "fmt"

type Trail_DataResource struct {
	
	
	Type interface{} `yaml:"Type"`
	Values interface{} `yaml:"Values,omitempty"`
}

func (resource Trail_DataResource) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
