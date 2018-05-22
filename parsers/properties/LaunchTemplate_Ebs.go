package properties


type LaunchTemplate_Ebs struct {
	
	
	
	
	
	
	
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	SnapshotId interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource LaunchTemplate_Ebs) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	return errs
}
