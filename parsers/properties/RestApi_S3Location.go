package properties


type RestApi_S3Location struct {
	
	
	
	
	Bucket interface{} `yaml:"Bucket,omitempty"`
	ETag interface{} `yaml:"ETag,omitempty"`
	Key interface{} `yaml:"Key,omitempty"`
	Version interface{} `yaml:"Version,omitempty"`
}

func (resource RestApi_S3Location) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
