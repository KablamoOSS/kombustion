package properties


type DeploymentGroup_AlarmConfiguration struct {
	
	
	
	Enabled interface{} `yaml:"Enabled,omitempty"`
	IgnorePollAlarmFailure interface{} `yaml:"IgnorePollAlarmFailure,omitempty"`
	Alarms interface{} `yaml:"Alarms,omitempty"`
}

func (resource DeploymentGroup_AlarmConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
