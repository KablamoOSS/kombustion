package properties


type Domain_VPCOptions struct {
	
	
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
	SubnetIds interface{} `yaml:"SubnetIds,omitempty"`
}

func (resource Domain_VPCOptions) Validate() []error {
	errs := []error{}
	
	
	return errs
}
