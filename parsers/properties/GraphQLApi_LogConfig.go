package properties


type GraphQLApi_LogConfig struct {
	
	
	CloudWatchLogsRoleArn interface{} `yaml:"CloudWatchLogsRoleArn,omitempty"`
	FieldLogLevel interface{} `yaml:"FieldLogLevel,omitempty"`
}

func (resource GraphQLApi_LogConfig) Validate() []error {
	errs := []error{}
	
	
	return errs
}
