package properties


type Layer_VolumeConfiguration struct {
	
	
	
	
	
	
	Iops interface{} `yaml:"Iops,omitempty"`
	MountPoint interface{} `yaml:"MountPoint,omitempty"`
	NumberOfDisks interface{} `yaml:"NumberOfDisks,omitempty"`
	RaidLevel interface{} `yaml:"RaidLevel,omitempty"`
	Size interface{} `yaml:"Size,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource Layer_VolumeConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
