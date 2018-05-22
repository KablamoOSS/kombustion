package properties


type Thing_AttributePayload struct {
	
	Attributes interface{} `yaml:"Attributes,omitempty"`
}

func (resource Thing_AttributePayload) Validate() []error {
	errs := []error{}
	
	return errs
}
