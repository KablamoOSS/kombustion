package properties


type Application_InputParallelism struct {
	
	Count interface{} `yaml:"Count,omitempty"`
}

func (resource Application_InputParallelism) Validate() []error {
	errs := []error{}
	
	return errs
}
