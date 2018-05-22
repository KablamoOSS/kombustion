package properties


type LoadBalancer_LBCookieStickinessPolicy struct {
	
	
	CookieExpirationPeriod interface{} `yaml:"CookieExpirationPeriod,omitempty"`
	PolicyName interface{} `yaml:"PolicyName,omitempty"`
}

func (resource LoadBalancer_LBCookieStickinessPolicy) Validate() []error {
	errs := []error{}
	
	
	return errs
}
