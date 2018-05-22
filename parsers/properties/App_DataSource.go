package properties


type App_DataSource struct {
	
	
	
	Arn interface{} `yaml:"Arn,omitempty"`
	DatabaseName interface{} `yaml:"DatabaseName,omitempty"`
	Type interface{} `yaml:"Type,omitempty"`
}

func (resource App_DataSource) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
