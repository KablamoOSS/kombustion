package properties


type JobDefinition_Volumes struct {
	
	
	Name interface{} `yaml:"Name,omitempty"`
	Host *JobDefinition_VolumesHost `yaml:"Host,omitempty"`
}

func (resource JobDefinition_Volumes) Validate() []error {
	errs := []error{}
	
	
	return errs
}
