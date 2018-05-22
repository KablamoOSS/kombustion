package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DMSEventSubscription struct {
	Type       string                      `yaml:"Type"`
	Properties DMSEventSubscriptionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DMSEventSubscriptionProperties struct {
	Enabled interface{} `yaml:"Enabled,omitempty"`
	SnsTopicArn interface{} `yaml:"SnsTopicArn"`
	SourceType interface{} `yaml:"SourceType,omitempty"`
	SubscriptionName interface{} `yaml:"SubscriptionName,omitempty"`
	EventCategories interface{} `yaml:"EventCategories,omitempty"`
	SourceIds interface{} `yaml:"SourceIds,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewDMSEventSubscription(properties DMSEventSubscriptionProperties, deps ...interface{}) DMSEventSubscription {
	return DMSEventSubscription{
		Type:       "AWS::DMS::EventSubscription",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDMSEventSubscription(name string, data string) (cf types.ValueMap, err error) {
	var resource DMSEventSubscription
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSEventSubscription - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DMSEventSubscription) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSEventSubscriptionProperties) Validate() []error {
	errs := []error{}
	if resource.SnsTopicArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SnsTopicArn'"))
	}
	return errs
}
