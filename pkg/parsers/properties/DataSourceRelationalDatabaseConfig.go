package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DataSourceRelationalDatabaseConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html
type DataSourceRelationalDatabaseConfig struct {
	RelationalDatabaseSourceType interface{} `yaml:"RelationalDatabaseSourceType"`
	RdsHttpEndpointConfig        interface{} `yaml:"RdsHttpEndpointConfig,omitempty"`
}

// DataSourceRelationalDatabaseConfig validation
func (resource DataSourceRelationalDatabaseConfig) Validate() []error {
	errors := []error{}

	return errors
}