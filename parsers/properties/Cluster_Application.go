package properties


type Cluster_Application struct {
	
	
	
	
	Name interface{} `yaml:"Name,omitempty"`
	Version interface{} `yaml:"Version,omitempty"`
	AdditionalInfo interface{} `yaml:"AdditionalInfo,omitempty"`
	Args interface{} `yaml:"Args,omitempty"`
}

func (resource Cluster_Application) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
