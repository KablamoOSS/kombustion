package properties


type LaunchTemplate_TagSpecification struct {
	
	
	ResourceType interface{} `yaml:"ResourceType,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func (resource LaunchTemplate_TagSpecification) Validate() []error {
	errs := []error{}
	
	
	return errs
}
