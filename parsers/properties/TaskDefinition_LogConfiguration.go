package properties

	import "fmt"

type TaskDefinition_LogConfiguration struct {
	
	
	LogDriver interface{} `yaml:"LogDriver"`
	Options interface{} `yaml:"Options,omitempty"`
}

func (resource TaskDefinition_LogConfiguration) Validate() []error {
	errs := []error{}
	
	
	if resource.LogDriver == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LogDriver'"))
	}
	return errs
}
