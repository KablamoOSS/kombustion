package properties

	import "fmt"

type JobQueue_ComputeEnvironmentOrder struct {
	
	
	ComputeEnvironment interface{} `yaml:"ComputeEnvironment"`
	Order interface{} `yaml:"Order"`
}

func (resource JobQueue_ComputeEnvironmentOrder) Validate() []error {
	errs := []error{}
	
	
	if resource.ComputeEnvironment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComputeEnvironment'"))
	}
	if resource.Order == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Order'"))
	}
	return errs
}
