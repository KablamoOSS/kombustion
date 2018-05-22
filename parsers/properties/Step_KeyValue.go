package properties


type Step_KeyValue struct {
	
	
	Key interface{} `yaml:"Key,omitempty"`
	Value interface{} `yaml:"Value,omitempty"`
}

func (resource Step_KeyValue) Validate() []error {
	errs := []error{}
	
	
	return errs
}
