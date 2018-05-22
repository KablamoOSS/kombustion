package properties

	import "fmt"

type DeploymentConfig_MinimumHealthyHosts struct {
	
	
	Type interface{} `yaml:"Type"`
	Value interface{} `yaml:"Value"`
}

func (resource DeploymentConfig_MinimumHealthyHosts) Validate() []error {
	errs := []error{}
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
