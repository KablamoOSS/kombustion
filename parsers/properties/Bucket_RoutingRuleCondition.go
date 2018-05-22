package properties


type Bucket_RoutingRuleCondition struct {
	
	
	HttpErrorCodeReturnedEquals interface{} `yaml:"HttpErrorCodeReturnedEquals,omitempty"`
	KeyPrefixEquals interface{} `yaml:"KeyPrefixEquals,omitempty"`
}

func (resource Bucket_RoutingRuleCondition) Validate() []error {
	errs := []error{}
	
	
	return errs
}
