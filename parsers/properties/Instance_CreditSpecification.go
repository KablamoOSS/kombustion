package properties


type Instance_CreditSpecification struct {
	
	CPUCredits interface{} `yaml:"CPUCredits,omitempty"`
}

func (resource Instance_CreditSpecification) Validate() []error {
	errs := []error{}
	
	return errs
}
