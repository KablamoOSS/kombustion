package properties


type DeploymentGroup_TriggerConfig struct {
	
	
	
	TriggerName interface{} `yaml:"TriggerName,omitempty"`
	TriggerTargetArn interface{} `yaml:"TriggerTargetArn,omitempty"`
	TriggerEvents interface{} `yaml:"TriggerEvents,omitempty"`
}

func (resource DeploymentGroup_TriggerConfig) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
