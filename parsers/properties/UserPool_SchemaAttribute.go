package properties


type UserPool_SchemaAttribute struct {
	
	
	
	
	
	
	
	AttributeDataType interface{} `yaml:"AttributeDataType,omitempty"`
	DeveloperOnlyAttribute interface{} `yaml:"DeveloperOnlyAttribute,omitempty"`
	Mutable interface{} `yaml:"Mutable,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Required interface{} `yaml:"Required,omitempty"`
	StringAttributeConstraints *UserPool_StringAttributeConstraints `yaml:"StringAttributeConstraints,omitempty"`
	NumberAttributeConstraints *UserPool_NumberAttributeConstraints `yaml:"NumberAttributeConstraints,omitempty"`
}

func (resource UserPool_SchemaAttribute) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	return errs
}
