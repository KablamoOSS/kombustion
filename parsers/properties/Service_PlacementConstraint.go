package properties

	import "fmt"

type Service_PlacementConstraint struct {
	
	
	Expression interface{} `yaml:"Expression,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource Service_PlacementConstraint) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
