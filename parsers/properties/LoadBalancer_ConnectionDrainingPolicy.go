package properties

	import "fmt"

type LoadBalancer_ConnectionDrainingPolicy struct {
	
	
	Enabled interface{} `yaml:"Enabled"`
	Timeout interface{} `yaml:"Timeout,omitempty"`
}

func (resource LoadBalancer_ConnectionDrainingPolicy) Validate() []error {
	errs := []error{}
	
	
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
