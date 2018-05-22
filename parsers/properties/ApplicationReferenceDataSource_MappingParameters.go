package properties


type ApplicationReferenceDataSource_MappingParameters struct {
	
	
	JSONMappingParameters *ApplicationReferenceDataSource_JSONMappingParameters `yaml:"JSONMappingParameters,omitempty"`
	CSVMappingParameters *ApplicationReferenceDataSource_CSVMappingParameters `yaml:"CSVMappingParameters,omitempty"`
}

func (resource ApplicationReferenceDataSource_MappingParameters) Validate() []error {
	errs := []error{}
	
	
	return errs
}
