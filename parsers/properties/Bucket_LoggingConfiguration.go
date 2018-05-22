package properties


type Bucket_LoggingConfiguration struct {
	
	
	DestinationBucketName interface{} `yaml:"DestinationBucketName,omitempty"`
	LogFilePrefix interface{} `yaml:"LogFilePrefix,omitempty"`
}

func (resource Bucket_LoggingConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
