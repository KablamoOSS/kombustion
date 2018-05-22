package properties

	import "fmt"

type HealthCheck_HealthCheckTag struct {
	
	
	Key interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

func (resource HealthCheck_HealthCheckTag) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
