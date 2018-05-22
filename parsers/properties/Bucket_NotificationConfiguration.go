package properties


type Bucket_NotificationConfiguration struct {
	
	
	
	LambdaConfigurations interface{} `yaml:"LambdaConfigurations,omitempty"`
	QueueConfigurations interface{} `yaml:"QueueConfigurations,omitempty"`
	TopicConfigurations interface{} `yaml:"TopicConfigurations,omitempty"`
}

func (resource Bucket_NotificationConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
