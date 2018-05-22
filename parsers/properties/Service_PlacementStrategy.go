package properties

	import "fmt"

type Service_PlacementStrategy struct {
	
	
	Field interface{} `yaml:"Field,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource Service_PlacementStrategy) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
