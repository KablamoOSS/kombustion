package properties


type Table_Projection struct {
	
	
	ProjectionType interface{} `yaml:"ProjectionType,omitempty"`
	NonKeyAttributes interface{} `yaml:"NonKeyAttributes,omitempty"`
}

func (resource Table_Projection) Validate() []error {
	errs := []error{}
	
	
	return errs
}
