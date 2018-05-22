package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type SNSTopic struct {
	Type       string                      `yaml:"Type"`
	Properties SNSTopicProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SNSTopicProperties struct {
	DisplayName interface{} `yaml:"DisplayName,omitempty"`
	TopicName interface{} `yaml:"TopicName,omitempty"`
	Subscription interface{} `yaml:"Subscription,omitempty"`
}

func NewSNSTopic(properties SNSTopicProperties, deps ...interface{}) SNSTopic {
	return SNSTopic{
		Type:       "AWS::SNS::Topic",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSNSTopic(name string, data string) (cf types.ValueMap, err error) {
	var resource SNSTopic
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SNSTopic - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SNSTopic) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SNSTopicProperties) Validate() []error {
	errs := []error{}
	return errs
}
