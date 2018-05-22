package properties


type LaunchConfiguration_BlockDevice struct {
	
	
	
	
	
	
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	SnapshotId interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource LaunchConfiguration_BlockDevice) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
