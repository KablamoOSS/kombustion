package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type LogsMetricFilter struct {
	Type       string                      `yaml:"Type"`
	Properties LogsMetricFilterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LogsMetricFilterProperties struct {
	FilterPattern interface{} `yaml:"FilterPattern"`
	LogGroupName interface{} `yaml:"LogGroupName"`
	MetricTransformations interface{} `yaml:"MetricTransformations"`
}

func NewLogsMetricFilter(properties LogsMetricFilterProperties, deps ...interface{}) LogsMetricFilter {
	return LogsMetricFilter{
		Type:       "AWS::Logs::MetricFilter",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLogsMetricFilter(name string, data string) (cf types.ValueMap, err error) {
	var resource LogsMetricFilter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsMetricFilter - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LogsMetricFilter) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LogsMetricFilterProperties) Validate() []error {
	errs := []error{}
	if resource.FilterPattern == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FilterPattern'"))
	}
	if resource.LogGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LogGroupName'"))
	}
	if resource.MetricTransformations == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricTransformations'"))
	}
	return errs
}
