package properties


type Application_InputProcessingConfiguration struct {
	
	InputLambdaProcessor *Application_InputLambdaProcessor `yaml:"InputLambdaProcessor,omitempty"`
}

func (resource Application_InputProcessingConfiguration) Validate() []error {
	errs := []error{}
	
	return errs
}
