package properties


type UsagePlan_QuotaSettings struct {
	
	
	
	Limit interface{} `yaml:"Limit,omitempty"`
	Offset interface{} `yaml:"Offset,omitempty"`
	Period interface{} `yaml:"Period,omitempty"`
}

func (resource UsagePlan_QuotaSettings) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
