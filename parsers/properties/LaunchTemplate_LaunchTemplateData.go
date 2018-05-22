package properties


type LaunchTemplate_LaunchTemplateData struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	DisableApiTermination interface{} `yaml:"DisableApiTermination,omitempty"`
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	ImageId interface{} `yaml:"ImageId,omitempty"`
	InstanceInitiatedShutdownBehavior interface{} `yaml:"InstanceInitiatedShutdownBehavior,omitempty"`
	InstanceType interface{} `yaml:"InstanceType,omitempty"`
	KernelId interface{} `yaml:"KernelId,omitempty"`
	KeyName interface{} `yaml:"KeyName,omitempty"`
	RamDiskId interface{} `yaml:"RamDiskId,omitempty"`
	UserData interface{} `yaml:"UserData,omitempty"`
	Placement *LaunchTemplate_Placement `yaml:"Placement,omitempty"`
	Monitoring *LaunchTemplate_Monitoring `yaml:"Monitoring,omitempty"`
	BlockDeviceMappings interface{} `yaml:"BlockDeviceMappings,omitempty"`
	ElasticGpuSpecifications interface{} `yaml:"ElasticGpuSpecifications,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	TagSpecifications interface{} `yaml:"TagSpecifications,omitempty"`
	NetworkInterfaces interface{} `yaml:"NetworkInterfaces,omitempty"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
	InstanceMarketOptions *LaunchTemplate_InstanceMarketOptions `yaml:"InstanceMarketOptions,omitempty"`
	IamInstanceProfile *LaunchTemplate_IamInstanceProfile `yaml:"IamInstanceProfile,omitempty"`
	CreditSpecification *LaunchTemplate_CreditSpecification `yaml:"CreditSpecification,omitempty"`
}

func (resource LaunchTemplate_LaunchTemplateData) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
