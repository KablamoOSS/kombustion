package properties


type Stack_Source struct {
	
	
	
	
	
	
	Password interface{} `yaml:"Password,omitempty"`
	Revision interface{} `yaml:"Revision,omitempty"`
	SshKey interface{} `yaml:"SshKey,omitempty"`
	Type interface{} `yaml:"Type,omitempty"`
	Url interface{} `yaml:"Url,omitempty"`
	Username interface{} `yaml:"Username,omitempty"`
}

func (resource Stack_Source) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
