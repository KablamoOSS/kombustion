package properties

	import "fmt"

type Project_Environment struct {
	
	
	
	
	
	ComputeType interface{} `yaml:"ComputeType"`
	Image interface{} `yaml:"Image"`
	PrivilegedMode interface{} `yaml:"PrivilegedMode,omitempty"`
	Type interface{} `yaml:"Type"`
	EnvironmentVariables interface{} `yaml:"EnvironmentVariables,omitempty"`
}

func (resource Project_Environment) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.ComputeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComputeType'"))
	}
	if resource.Image == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Image'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
