package properties

	import "fmt"

type LoadBalancer_ConnectionSettings struct {
	
	IdleTimeout interface{} `yaml:"IdleTimeout"`
}

func (resource LoadBalancer_ConnectionSettings) Validate() []error {
	errs := []error{}
	
	if resource.IdleTimeout == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IdleTimeout'"))
	}
	return errs
}
