package properties

	import "fmt"

type LoadBalancer_HealthCheck struct {
	
	
	
	
	
	HealthyThreshold interface{} `yaml:"HealthyThreshold"`
	Interval interface{} `yaml:"Interval"`
	Target interface{} `yaml:"Target"`
	Timeout interface{} `yaml:"Timeout"`
	UnhealthyThreshold interface{} `yaml:"UnhealthyThreshold"`
}

func (resource LoadBalancer_HealthCheck) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.HealthyThreshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HealthyThreshold'"))
	}
	if resource.Interval == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Interval'"))
	}
	if resource.Target == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Target'"))
	}
	if resource.Timeout == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Timeout'"))
	}
	if resource.UnhealthyThreshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UnhealthyThreshold'"))
	}
	return errs
}
