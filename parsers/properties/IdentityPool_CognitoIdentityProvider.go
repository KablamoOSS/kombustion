package properties


type IdentityPool_CognitoIdentityProvider struct {
	
	
	
	ClientId interface{} `yaml:"ClientId,omitempty"`
	ProviderName interface{} `yaml:"ProviderName,omitempty"`
	ServerSideTokenCheck interface{} `yaml:"ServerSideTokenCheck,omitempty"`
}

func (resource IdentityPool_CognitoIdentityProvider) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
