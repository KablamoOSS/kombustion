package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type IoTTopicRule struct {
	Type       string                      `yaml:"Type"`
	Properties IoTTopicRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IoTTopicRuleProperties struct {
	RuleName interface{} `yaml:"RuleName,omitempty"`
	TopicRulePayload *properties.TopicRule_TopicRulePayload `yaml:"TopicRulePayload"`
}

func NewIoTTopicRule(properties IoTTopicRuleProperties, deps ...interface{}) IoTTopicRule {
	return IoTTopicRule{
		Type:       "AWS::IoT::TopicRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTTopicRule(name string, data string) (cf types.ValueMap, err error) {
	var resource IoTTopicRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTTopicRule - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IoTTopicRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTTopicRuleProperties) Validate() []error {
	errs := []error{}
	if resource.TopicRulePayload == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TopicRulePayload'"))
	} else {
		errs = append(errs, resource.TopicRulePayload.Validate()...)
	}
	return errs
}
