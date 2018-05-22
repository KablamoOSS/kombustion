package properties


type Job_ConnectionsList struct {
	
	Connections interface{} `yaml:"Connections,omitempty"`
}

func (resource Job_ConnectionsList) Validate() []error {
	errs := []error{}
	
	return errs
}
