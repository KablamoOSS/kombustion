package properties


type Cluster_Configuration struct {
	
	
	
	Classification interface{} `yaml:"Classification,omitempty"`
	ConfigurationProperties interface{} `yaml:"ConfigurationProperties,omitempty"`
	Configurations interface{} `yaml:"Configurations,omitempty"`
}

func (resource Cluster_Configuration) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
