package properties


type Function_Code struct {
	
	
	
	
	S3Bucket interface{} `yaml:"S3Bucket,omitempty"`
	S3Key interface{} `yaml:"S3Key,omitempty"`
	S3ObjectVersion interface{} `yaml:"S3ObjectVersion,omitempty"`
	ZipFile interface{} `yaml:"ZipFile,omitempty"`
}

func (resource Function_Code) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
