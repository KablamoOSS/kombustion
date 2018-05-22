package properties


type Instance_Ebs struct {
	
	
	
	
	
	
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	SnapshotId interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource Instance_Ebs) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
