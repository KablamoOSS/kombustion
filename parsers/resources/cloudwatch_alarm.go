package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CloudWatchAlarm struct {
	Type       string                      `yaml:"Type"`
	Properties CloudWatchAlarmProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudWatchAlarmProperties struct {
	ActionsEnabled interface{} `yaml:"ActionsEnabled,omitempty"`
	AlarmDescription interface{} `yaml:"AlarmDescription,omitempty"`
	AlarmName interface{} `yaml:"AlarmName,omitempty"`
	ComparisonOperator interface{} `yaml:"ComparisonOperator"`
	EvaluateLowSampleCountPercentile interface{} `yaml:"EvaluateLowSampleCountPercentile,omitempty"`
	EvaluationPeriods interface{} `yaml:"EvaluationPeriods"`
	ExtendedStatistic interface{} `yaml:"ExtendedStatistic,omitempty"`
	MetricName interface{} `yaml:"MetricName"`
	Namespace interface{} `yaml:"Namespace"`
	Period interface{} `yaml:"Period"`
	Statistic interface{} `yaml:"Statistic,omitempty"`
	Threshold interface{} `yaml:"Threshold"`
	TreatMissingData interface{} `yaml:"TreatMissingData,omitempty"`
	Unit interface{} `yaml:"Unit,omitempty"`
	AlarmActions interface{} `yaml:"AlarmActions,omitempty"`
	Dimensions interface{} `yaml:"Dimensions,omitempty"`
	InsufficientDataActions interface{} `yaml:"InsufficientDataActions,omitempty"`
	OKActions interface{} `yaml:"OKActions,omitempty"`
}

func NewCloudWatchAlarm(properties CloudWatchAlarmProperties, deps ...interface{}) CloudWatchAlarm {
	return CloudWatchAlarm{
		Type:       "AWS::CloudWatch::Alarm",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudWatchAlarm(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudWatchAlarm
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudWatchAlarm - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudWatchAlarm) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudWatchAlarmProperties) Validate() []error {
	errs := []error{}
	if resource.ComparisonOperator == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComparisonOperator'"))
	}
	if resource.EvaluationPeriods == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EvaluationPeriods'"))
	}
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Namespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Namespace'"))
	}
	if resource.Period == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Period'"))
	}
	if resource.Threshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Threshold'"))
	}
	return errs
}
