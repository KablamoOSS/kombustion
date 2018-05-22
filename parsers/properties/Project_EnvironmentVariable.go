package properties

	import "fmt"

type Project_EnvironmentVariable struct {
	
	
	
	Name interface{} `yaml:"Name"`
	Type interface{} `yaml:"Type,omitempty"`
	Value interface{} `yaml:"Value"`
}

func (resource Project_EnvironmentVariable) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
