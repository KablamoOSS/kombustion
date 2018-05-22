package properties


type DBSecurityGroup_Ingress struct {
	
	
	
	
	CIDRIP interface{} `yaml:"CIDRIP,omitempty"`
	EC2SecurityGroupId interface{} `yaml:"EC2SecurityGroupId,omitempty"`
	EC2SecurityGroupName interface{} `yaml:"EC2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerId interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

func (resource DBSecurityGroup_Ingress) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
