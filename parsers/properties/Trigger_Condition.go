package properties


type Trigger_Condition struct {
	
	
	
	JobName interface{} `yaml:"JobName,omitempty"`
	LogicalOperator interface{} `yaml:"LogicalOperator,omitempty"`
	State interface{} `yaml:"State,omitempty"`
}

func (resource Trigger_Condition) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
