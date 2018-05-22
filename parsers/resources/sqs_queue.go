package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type SQSQueue struct {
	Type       string                      `yaml:"Type"`
	Properties SQSQueueProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SQSQueueProperties struct {
	ContentBasedDeduplication interface{} `yaml:"ContentBasedDeduplication,omitempty"`
	DelaySeconds interface{} `yaml:"DelaySeconds,omitempty"`
	FifoQueue interface{} `yaml:"FifoQueue,omitempty"`
	KmsDataKeyReusePeriodSeconds interface{} `yaml:"KmsDataKeyReusePeriodSeconds,omitempty"`
	KmsMasterKeyId interface{} `yaml:"KmsMasterKeyId,omitempty"`
	MaximumMessageSize interface{} `yaml:"MaximumMessageSize,omitempty"`
	MessageRetentionPeriod interface{} `yaml:"MessageRetentionPeriod,omitempty"`
	QueueName interface{} `yaml:"QueueName,omitempty"`
	ReceiveMessageWaitTimeSeconds interface{} `yaml:"ReceiveMessageWaitTimeSeconds,omitempty"`
	RedrivePolicy interface{} `yaml:"RedrivePolicy,omitempty"`
	VisibilityTimeout interface{} `yaml:"VisibilityTimeout,omitempty"`
}

func NewSQSQueue(properties SQSQueueProperties, deps ...interface{}) SQSQueue {
	return SQSQueue{
		Type:       "AWS::SQS::Queue",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSQSQueue(name string, data string) (cf types.ValueMap, err error) {
	var resource SQSQueue
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SQSQueue - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SQSQueue) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SQSQueueProperties) Validate() []error {
	errs := []error{}
	return errs
}
