package properties


type UserPool_StringAttributeConstraints struct {
	
	
	MaxLength interface{} `yaml:"MaxLength,omitempty"`
	MinLength interface{} `yaml:"MinLength,omitempty"`
}

func (resource UserPool_StringAttributeConstraints) Validate() []error {
	errs := []error{}
	
	
	return errs
}
