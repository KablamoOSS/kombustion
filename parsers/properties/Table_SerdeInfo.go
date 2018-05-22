package properties


type Table_SerdeInfo struct {
	
	
	
	Name interface{} `yaml:"Name,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	SerializationLibrary interface{} `yaml:"SerializationLibrary,omitempty"`
}

func (resource Table_SerdeInfo) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
