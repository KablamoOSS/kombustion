package properties


type TaskDefinition_PortMapping struct {
	
	
	
	ContainerPort interface{} `yaml:"ContainerPort,omitempty"`
	HostPort interface{} `yaml:"HostPort,omitempty"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
}

func (resource TaskDefinition_PortMapping) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
