package properties


type TopicRule_DynamoDBv2Action struct {
	
	
	RoleArn interface{} `yaml:"RoleArn,omitempty"`
	PutItem *TopicRule_PutItemInput `yaml:"PutItem,omitempty"`
}

func (resource TopicRule_DynamoDBv2Action) Validate() []error {
	errs := []error{}
	
	
	return errs
}
