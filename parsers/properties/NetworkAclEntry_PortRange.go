package properties


type NetworkAclEntry_PortRange struct {
	
	
	From interface{} `yaml:"From,omitempty"`
	To interface{} `yaml:"To,omitempty"`
}

func (resource NetworkAclEntry_PortRange) Validate() []error {
	errs := []error{}
	
	
	return errs
}
