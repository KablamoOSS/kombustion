package properties


type SpotFleet_EbsBlockDevice struct {
	
	
	
	
	
	
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	SnapshotId interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource SpotFleet_EbsBlockDevice) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
