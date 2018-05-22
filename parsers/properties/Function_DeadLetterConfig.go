package properties


type Function_DeadLetterConfig struct {
	
	TargetArn interface{} `yaml:"TargetArn,omitempty"`
}

func (resource Function_DeadLetterConfig) Validate() []error {
	errs := []error{}
	
	return errs
}
