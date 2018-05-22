package properties


type Job_JobCommand struct {
	
	
	Name interface{} `yaml:"Name,omitempty"`
	ScriptLocation interface{} `yaml:"ScriptLocation,omitempty"`
}

func (resource Job_JobCommand) Validate() []error {
	errs := []error{}
	
	
	return errs
}
