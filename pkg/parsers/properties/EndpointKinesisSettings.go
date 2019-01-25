package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EndpointKinesisSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html
type EndpointKinesisSettings struct {
	MessageFormat        interface{} `yaml:"MessageFormat,omitempty"`
	ServiceAccessRoleArn interface{} `yaml:"ServiceAccessRoleArn,omitempty"`
	StreamArn            interface{} `yaml:"StreamArn,omitempty"`
}

// EndpointKinesisSettings validation
func (resource EndpointKinesisSettings) Validate() []error {
	errors := []error{}

	return errors
}
