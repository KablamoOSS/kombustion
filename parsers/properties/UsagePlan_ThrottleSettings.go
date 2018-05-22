package properties


type UsagePlan_ThrottleSettings struct {
	
	
	BurstLimit interface{} `yaml:"BurstLimit,omitempty"`
	RateLimit interface{} `yaml:"RateLimit,omitempty"`
}

func (resource UsagePlan_ThrottleSettings) Validate() []error {
	errs := []error{}
	
	
	return errs
}
