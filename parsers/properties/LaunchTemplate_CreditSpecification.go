package properties


type LaunchTemplate_CreditSpecification struct {
	
	CpuCredits interface{} `yaml:"CpuCredits,omitempty"`
}

func (resource LaunchTemplate_CreditSpecification) Validate() []error {
	errs := []error{}
	
	return errs
}
