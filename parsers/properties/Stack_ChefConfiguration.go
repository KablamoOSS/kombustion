package properties


type Stack_ChefConfiguration struct {
	
	
	BerkshelfVersion interface{} `yaml:"BerkshelfVersion,omitempty"`
	ManageBerkshelf interface{} `yaml:"ManageBerkshelf,omitempty"`
}

func (resource Stack_ChefConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
