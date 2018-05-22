package properties


type App_SslConfiguration struct {
	
	
	
	Certificate interface{} `yaml:"Certificate,omitempty"`
	Chain interface{} `yaml:"Chain,omitempty"`
	PrivateKey interface{} `yaml:"PrivateKey,omitempty"`
}

func (resource App_SslConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
