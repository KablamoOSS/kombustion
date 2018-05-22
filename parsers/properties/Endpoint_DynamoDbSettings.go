package properties


type Endpoint_DynamoDbSettings struct {
	
	ServiceAccessRoleArn interface{} `yaml:"ServiceAccessRoleArn,omitempty"`
}

func (resource Endpoint_DynamoDbSettings) Validate() []error {
	errs := []error{}
	
	return errs
}
