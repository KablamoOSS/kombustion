package properties

	import "fmt"

type Rule_RunCommandParameters struct {
	
	RunCommandTargets interface{} `yaml:"RunCommandTargets"`
}

func (resource Rule_RunCommandParameters) Validate() []error {
	errs := []error{}
	
	if resource.RunCommandTargets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RunCommandTargets'"))
	}
	return errs
}
