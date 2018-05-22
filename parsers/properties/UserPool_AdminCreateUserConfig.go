package properties


type UserPool_AdminCreateUserConfig struct {
	
	
	
	AllowAdminCreateUserOnly interface{} `yaml:"AllowAdminCreateUserOnly,omitempty"`
	UnusedAccountValidityDays interface{} `yaml:"UnusedAccountValidityDays,omitempty"`
	InviteMessageTemplate *UserPool_InviteMessageTemplate `yaml:"InviteMessageTemplate,omitempty"`
}

func (resource UserPool_AdminCreateUserConfig) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
