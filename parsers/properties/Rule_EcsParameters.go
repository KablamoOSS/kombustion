package properties

	import "fmt"

type Rule_EcsParameters struct {
	
	
	TaskCount interface{} `yaml:"TaskCount,omitempty"`
	TaskDefinitionArn interface{} `yaml:"TaskDefinitionArn"`
}

func (resource Rule_EcsParameters) Validate() []error {
	errs := []error{}
	
	
	if resource.TaskDefinitionArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TaskDefinitionArn'"))
	}
	return errs
}
