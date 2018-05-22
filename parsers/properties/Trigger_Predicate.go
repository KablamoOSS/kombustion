package properties


type Trigger_Predicate struct {
	
	
	Logical interface{} `yaml:"Logical,omitempty"`
	Conditions interface{} `yaml:"Conditions,omitempty"`
}

func (resource Trigger_Predicate) Validate() []error {
	errs := []error{}
	
	
	return errs
}
