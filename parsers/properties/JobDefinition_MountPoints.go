package properties


type JobDefinition_MountPoints struct {
	
	
	
	ContainerPath interface{} `yaml:"ContainerPath,omitempty"`
	ReadOnly interface{} `yaml:"ReadOnly,omitempty"`
	SourceVolume interface{} `yaml:"SourceVolume,omitempty"`
}

func (resource JobDefinition_MountPoints) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
