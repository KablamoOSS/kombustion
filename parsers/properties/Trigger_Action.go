package properties


type Trigger_Action struct {
	
	
	Arguments interface{} `yaml:"Arguments,omitempty"`
	JobName interface{} `yaml:"JobName,omitempty"`
}

func (resource Trigger_Action) Validate() []error {
	errs := []error{}
	
	
	return errs
}
