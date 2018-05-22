package properties


type TopicRule_Action struct {
	
	
	
	
	
	
	
	
	
	
	
	
	Sqs *TopicRule_SqsAction `yaml:"Sqs,omitempty"`
	Sns *TopicRule_SnsAction `yaml:"Sns,omitempty"`
	S3 *TopicRule_S3Action `yaml:"S3,omitempty"`
	Republish *TopicRule_RepublishAction `yaml:"Republish,omitempty"`
	Lambda *TopicRule_LambdaAction `yaml:"Lambda,omitempty"`
	Kinesis *TopicRule_KinesisAction `yaml:"Kinesis,omitempty"`
	Firehose *TopicRule_FirehoseAction `yaml:"Firehose,omitempty"`
	Elasticsearch *TopicRule_ElasticsearchAction `yaml:"Elasticsearch,omitempty"`
	DynamoDBv2 *TopicRule_DynamoDBv2Action `yaml:"DynamoDBv2,omitempty"`
	DynamoDB *TopicRule_DynamoDBAction `yaml:"DynamoDB,omitempty"`
	CloudwatchMetric *TopicRule_CloudwatchMetricAction `yaml:"CloudwatchMetric,omitempty"`
	CloudwatchAlarm *TopicRule_CloudwatchAlarmAction `yaml:"CloudwatchAlarm,omitempty"`
}

func (resource TopicRule_Action) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
