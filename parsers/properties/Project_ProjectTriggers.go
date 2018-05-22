package properties


type Project_ProjectTriggers struct {
	
	Webhook interface{} `yaml:"Webhook,omitempty"`
}

func (resource Project_ProjectTriggers) Validate() []error {
	errs := []error{}
	
	return errs
}
