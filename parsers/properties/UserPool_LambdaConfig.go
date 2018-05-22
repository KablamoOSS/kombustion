package properties


type UserPool_LambdaConfig struct {
	
	
	
	
	
	
	
	
	CreateAuthChallenge interface{} `yaml:"CreateAuthChallenge,omitempty"`
	CustomMessage interface{} `yaml:"CustomMessage,omitempty"`
	DefineAuthChallenge interface{} `yaml:"DefineAuthChallenge,omitempty"`
	PostAuthentication interface{} `yaml:"PostAuthentication,omitempty"`
	PostConfirmation interface{} `yaml:"PostConfirmation,omitempty"`
	PreAuthentication interface{} `yaml:"PreAuthentication,omitempty"`
	PreSignUp interface{} `yaml:"PreSignUp,omitempty"`
	VerifyAuthChallengeResponse interface{} `yaml:"VerifyAuthChallengeResponse,omitempty"`
}

func (resource UserPool_LambdaConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	return errs
}
