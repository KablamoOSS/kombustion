package properties


type Repository_RepositoryTrigger struct {
	
	
	
	
	
	CustomData interface{} `yaml:"CustomData,omitempty"`
	DestinationArn interface{} `yaml:"DestinationArn,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Branches interface{} `yaml:"Branches,omitempty"`
	Events interface{} `yaml:"Events,omitempty"`
}

func (resource Repository_RepositoryTrigger) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
