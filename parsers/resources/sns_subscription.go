package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type SNSSubscription struct {
	Type       string                      `yaml:"Type"`
	Properties SNSSubscriptionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SNSSubscriptionProperties struct {
	Endpoint interface{} `yaml:"Endpoint,omitempty"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

func NewSNSSubscription(properties SNSSubscriptionProperties, deps ...interface{}) SNSSubscription {
	return SNSSubscription{
		Type:       "AWS::SNS::Subscription",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSNSSubscription(name string, data string) (cf types.ValueMap, err error) {
	var resource SNSSubscription
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SNSSubscription - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SNSSubscription) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SNSSubscriptionProperties) Validate() []error {
	errs := []error{}
	return errs
}
