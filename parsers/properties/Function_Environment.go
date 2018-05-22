package properties


type Function_Environment struct {
	
	Variables interface{} `yaml:"Variables,omitempty"`
}

func (resource Function_Environment) Validate() []error {
	errs := []error{}
	
	return errs
}
