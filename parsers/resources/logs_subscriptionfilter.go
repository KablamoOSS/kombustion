package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type LogsSubscriptionFilter struct {
	Type       string                      `yaml:"Type"`
	Properties LogsSubscriptionFilterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LogsSubscriptionFilterProperties struct {
	DestinationArn interface{} `yaml:"DestinationArn"`
	FilterPattern interface{} `yaml:"FilterPattern"`
	LogGroupName interface{} `yaml:"LogGroupName"`
	RoleArn interface{} `yaml:"RoleArn,omitempty"`
}

func NewLogsSubscriptionFilter(properties LogsSubscriptionFilterProperties, deps ...interface{}) LogsSubscriptionFilter {
	return LogsSubscriptionFilter{
		Type:       "AWS::Logs::SubscriptionFilter",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLogsSubscriptionFilter(name string, data string) (cf types.ValueMap, err error) {
	var resource LogsSubscriptionFilter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsSubscriptionFilter - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LogsSubscriptionFilter) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LogsSubscriptionFilterProperties) Validate() []error {
	errs := []error{}
	if resource.DestinationArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DestinationArn'"))
	}
	if resource.FilterPattern == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FilterPattern'"))
	}
	if resource.LogGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LogGroupName'"))
	}
	return errs
}
