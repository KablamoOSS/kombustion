package properties

	import "fmt"

type Project_ProjectCache struct {
	
	
	Location interface{} `yaml:"Location,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource Project_ProjectCache) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
