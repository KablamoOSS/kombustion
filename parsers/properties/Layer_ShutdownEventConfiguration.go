package properties


type Layer_ShutdownEventConfiguration struct {
	
	
	DelayUntilElbConnectionsDrained interface{} `yaml:"DelayUntilElbConnectionsDrained,omitempty"`
	ExecutionTimeout interface{} `yaml:"ExecutionTimeout,omitempty"`
}

func (resource Layer_ShutdownEventConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
