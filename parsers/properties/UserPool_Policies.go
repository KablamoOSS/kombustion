package properties


type UserPool_Policies struct {
	
	PasswordPolicy *UserPool_PasswordPolicy `yaml:"PasswordPolicy,omitempty"`
}

func (resource UserPool_Policies) Validate() []error {
	errs := []error{}
	
	return errs
}
