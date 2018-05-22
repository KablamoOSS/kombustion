package properties


type LaunchTemplate_BlockDeviceMapping struct {
	
	
	
	
	DeviceName interface{} `yaml:"DeviceName,omitempty"`
	NoDevice interface{} `yaml:"NoDevice,omitempty"`
	VirtualName interface{} `yaml:"VirtualName,omitempty"`
	Ebs *LaunchTemplate_Ebs `yaml:"Ebs,omitempty"`
}

func (resource LaunchTemplate_BlockDeviceMapping) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
