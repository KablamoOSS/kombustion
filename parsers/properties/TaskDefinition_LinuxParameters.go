package properties


type TaskDefinition_LinuxParameters struct {
	
	
	
	InitProcessEnabled interface{} `yaml:"InitProcessEnabled,omitempty"`
	Devices interface{} `yaml:"Devices,omitempty"`
	Capabilities *TaskDefinition_KernelCapabilities `yaml:"Capabilities,omitempty"`
}

func (resource TaskDefinition_LinuxParameters) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
