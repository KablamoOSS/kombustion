package properties


type JobDefinition_RetryStrategy struct {
	
	Attempts interface{} `yaml:"Attempts,omitempty"`
}

func (resource JobDefinition_RetryStrategy) Validate() []error {
	errs := []error{}
	
	return errs
}
