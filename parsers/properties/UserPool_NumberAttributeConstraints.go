package properties


type UserPool_NumberAttributeConstraints struct {
	
	
	MaxValue interface{} `yaml:"MaxValue,omitempty"`
	MinValue interface{} `yaml:"MinValue,omitempty"`
}

func (resource UserPool_NumberAttributeConstraints) Validate() []error {
	errs := []error{}
	
	
	return errs
}
