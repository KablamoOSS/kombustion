package properties


type Layer_Recipes struct {
	
	
	
	
	
	Configure interface{} `yaml:"Configure,omitempty"`
	Deploy interface{} `yaml:"Deploy,omitempty"`
	Setup interface{} `yaml:"Setup,omitempty"`
	Shutdown interface{} `yaml:"Shutdown,omitempty"`
	Undeploy interface{} `yaml:"Undeploy,omitempty"`
}

func (resource Layer_Recipes) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
