package properties


type UserPool_PasswordPolicy struct {
	
	
	
	
	
	MinimumLength interface{} `yaml:"MinimumLength,omitempty"`
	RequireLowercase interface{} `yaml:"RequireLowercase,omitempty"`
	RequireNumbers interface{} `yaml:"RequireNumbers,omitempty"`
	RequireSymbols interface{} `yaml:"RequireSymbols,omitempty"`
	RequireUppercase interface{} `yaml:"RequireUppercase,omitempty"`
}

func (resource UserPool_PasswordPolicy) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
