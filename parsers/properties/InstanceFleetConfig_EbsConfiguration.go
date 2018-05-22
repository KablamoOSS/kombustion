package properties


type InstanceFleetConfig_EbsConfiguration struct {
	
	
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	EbsBlockDeviceConfigs interface{} `yaml:"EbsBlockDeviceConfigs,omitempty"`
}

func (resource InstanceFleetConfig_EbsConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
