package properties


type Layer_LifecycleEventConfiguration struct {
	
	ShutdownEventConfiguration *Layer_ShutdownEventConfiguration `yaml:"ShutdownEventConfiguration,omitempty"`
}

func (resource Layer_LifecycleEventConfiguration) Validate() []error {
	errs := []error{}
	
	return errs
}
