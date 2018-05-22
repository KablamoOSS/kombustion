package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type SQSQueuePolicy struct {
	Type       string                      `yaml:"Type"`
	Properties SQSQueuePolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SQSQueuePolicyProperties struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	Queues interface{} `yaml:"Queues"`
}

func NewSQSQueuePolicy(properties SQSQueuePolicyProperties, deps ...interface{}) SQSQueuePolicy {
	return SQSQueuePolicy{
		Type:       "AWS::SQS::QueuePolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSQSQueuePolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource SQSQueuePolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SQSQueuePolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SQSQueuePolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SQSQueuePolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.Queues == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Queues'"))
	}
	return errs
}
