package properties

	import "fmt"

type App_EnvironmentVariable struct {
	
	
	
	Key interface{} `yaml:"Key"`
	Secure interface{} `yaml:"Secure,omitempty"`
	Value interface{} `yaml:"Value"`
}

func (resource App_EnvironmentVariable) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
