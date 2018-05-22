package properties


type Domain_ElasticsearchClusterConfig struct {
	
	
	
	
	
	
	DedicatedMasterCount interface{} `yaml:"DedicatedMasterCount,omitempty"`
	DedicatedMasterEnabled interface{} `yaml:"DedicatedMasterEnabled,omitempty"`
	DedicatedMasterType interface{} `yaml:"DedicatedMasterType,omitempty"`
	InstanceCount interface{} `yaml:"InstanceCount,omitempty"`
	InstanceType interface{} `yaml:"InstanceType,omitempty"`
	ZoneAwarenessEnabled interface{} `yaml:"ZoneAwarenessEnabled,omitempty"`
}

func (resource Domain_ElasticsearchClusterConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
