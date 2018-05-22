package properties


type Endpoint_S3Settings struct {
	
	
	
	
	
	
	
	BucketFolder interface{} `yaml:"BucketFolder,omitempty"`
	BucketName interface{} `yaml:"BucketName,omitempty"`
	CompressionType interface{} `yaml:"CompressionType,omitempty"`
	CsvDelimiter interface{} `yaml:"CsvDelimiter,omitempty"`
	CsvRowDelimiter interface{} `yaml:"CsvRowDelimiter,omitempty"`
	ExternalTableDefinition interface{} `yaml:"ExternalTableDefinition,omitempty"`
	ServiceAccessRoleArn interface{} `yaml:"ServiceAccessRoleArn,omitempty"`
}

func (resource Endpoint_S3Settings) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	return errs
}
