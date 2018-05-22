package properties


type InstanceGroupConfig_EbsConfiguration struct {
	
	
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	EbsBlockDeviceConfigs interface{} `yaml:"EbsBlockDeviceConfigs,omitempty"`
}

func (resource InstanceGroupConfig_EbsConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
