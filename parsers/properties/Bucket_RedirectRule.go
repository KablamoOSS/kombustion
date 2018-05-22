package properties


type Bucket_RedirectRule struct {
	
	
	
	
	
	HostName interface{} `yaml:"HostName,omitempty"`
	HttpRedirectCode interface{} `yaml:"HttpRedirectCode,omitempty"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
	ReplaceKeyPrefixWith interface{} `yaml:"ReplaceKeyPrefixWith,omitempty"`
	ReplaceKeyWith interface{} `yaml:"ReplaceKeyWith,omitempty"`
}

func (resource Bucket_RedirectRule) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
