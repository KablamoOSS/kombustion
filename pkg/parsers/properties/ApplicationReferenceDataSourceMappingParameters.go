package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationReferenceDataSourceMappingParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-mappingparameters.html
type ApplicationReferenceDataSourceMappingParameters struct {
	JSONMappingParameters *ApplicationReferenceDataSourceJSONMappingParameters `yaml:"JSONMappingParameters,omitempty"`
	CSVMappingParameters  *ApplicationReferenceDataSourceCSVMappingParameters  `yaml:"CSVMappingParameters,omitempty"`
}

// ApplicationReferenceDataSourceMappingParameters validation
func (resource ApplicationReferenceDataSourceMappingParameters) Validate() []error {
	errors := []error{}

	return errors
}
