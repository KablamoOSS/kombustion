package properties


type Instance_EbsBlockDevice struct {
	
	
	
	
	
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	SnapshotId interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource Instance_EbsBlockDevice) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
