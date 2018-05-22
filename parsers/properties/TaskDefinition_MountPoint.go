package properties


type TaskDefinition_MountPoint struct {
	
	
	
	ContainerPath interface{} `yaml:"ContainerPath,omitempty"`
	ReadOnly interface{} `yaml:"ReadOnly,omitempty"`
	SourceVolume interface{} `yaml:"SourceVolume,omitempty"`
}

func (resource TaskDefinition_MountPoint) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
