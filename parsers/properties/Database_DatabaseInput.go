package properties


type Database_DatabaseInput struct {
	
	
	
	
	Description interface{} `yaml:"Description,omitempty"`
	LocationUri interface{} `yaml:"LocationUri,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
}

func (resource Database_DatabaseInput) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
