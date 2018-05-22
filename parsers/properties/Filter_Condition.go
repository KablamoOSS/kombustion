package properties


type Filter_Condition struct {
	
	
	
	
	
	Gte interface{} `yaml:"Gte,omitempty"`
	Lt interface{} `yaml:"Lt,omitempty"`
	Lte interface{} `yaml:"Lte,omitempty"`
	Eq interface{} `yaml:"Eq,omitempty"`
	Neq interface{} `yaml:"Neq,omitempty"`
}

func (resource Filter_Condition) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
