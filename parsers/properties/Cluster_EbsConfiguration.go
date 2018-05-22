package properties


type Cluster_EbsConfiguration struct {
	
	
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	EbsBlockDeviceConfigs interface{} `yaml:"EbsBlockDeviceConfigs,omitempty"`
}

func (resource Cluster_EbsConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
