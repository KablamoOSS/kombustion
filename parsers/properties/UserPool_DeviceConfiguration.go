package properties


type UserPool_DeviceConfiguration struct {
	
	
	ChallengeRequiredOnNewDevice interface{} `yaml:"ChallengeRequiredOnNewDevice,omitempty"`
	DeviceOnlyRememberedOnUserPrompt interface{} `yaml:"DeviceOnlyRememberedOnUserPrompt,omitempty"`
}

func (resource UserPool_DeviceConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
