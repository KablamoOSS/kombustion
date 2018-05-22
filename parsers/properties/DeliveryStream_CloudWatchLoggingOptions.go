package properties


type DeliveryStream_CloudWatchLoggingOptions struct {
	
	
	
	Enabled interface{} `yaml:"Enabled,omitempty"`
	LogGroupName interface{} `yaml:"LogGroupName,omitempty"`
	LogStreamName interface{} `yaml:"LogStreamName,omitempty"`
}

func (resource DeliveryStream_CloudWatchLoggingOptions) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
