package properties


type Service_DeploymentConfiguration struct {
	
	
	MaximumPercent interface{} `yaml:"MaximumPercent,omitempty"`
	MinimumHealthyPercent interface{} `yaml:"MinimumHealthyPercent,omitempty"`
}

func (resource Service_DeploymentConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
