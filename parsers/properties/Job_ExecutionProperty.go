package properties


type Job_ExecutionProperty struct {
	
	MaxConcurrentRuns interface{} `yaml:"MaxConcurrentRuns,omitempty"`
}

func (resource Job_ExecutionProperty) Validate() []error {
	errs := []error{}
	
	return errs
}
