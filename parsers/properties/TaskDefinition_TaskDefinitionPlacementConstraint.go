package properties

	import "fmt"

type TaskDefinition_TaskDefinitionPlacementConstraint struct {
	
	
	Expression interface{} `yaml:"Expression,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource TaskDefinition_TaskDefinitionPlacementConstraint) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
