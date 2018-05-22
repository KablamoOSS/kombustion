package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type SNSTopicPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties SNSTopicPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SNSTopicPolicyProperties struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	Topics interface{} `yaml:"Topics"`
}

func NewSNSTopicPolicy(properties SNSTopicPolicyProperties, deps ...interface{}) SNSTopicPolicy {
	return SNSTopicPolicy{
		Type:       "AWS::SNS::TopicPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSNSTopicPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource SNSTopicPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SNSTopicPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SNSTopicPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SNSTopicPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.Topics == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Topics'"))
	}
	return errs
}
