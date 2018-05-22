package properties


type TaskDefinition_VolumeFrom struct {
	
	
	ReadOnly interface{} `yaml:"ReadOnly,omitempty"`
	SourceContainer interface{} `yaml:"SourceContainer,omitempty"`
}

func (resource TaskDefinition_VolumeFrom) Validate() []error {
	errs := []error{}
	
	
	return errs
}
