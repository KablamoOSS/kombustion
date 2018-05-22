package properties


type TaskDefinition_Volume struct {
	
	
	Name interface{} `yaml:"Name,omitempty"`
	Host *TaskDefinition_HostVolumeProperties `yaml:"Host,omitempty"`
}

func (resource TaskDefinition_Volume) Validate() []error {
	errs := []error{}
	
	
	return errs
}
