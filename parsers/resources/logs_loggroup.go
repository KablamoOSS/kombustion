package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type LogsLogGroup struct {
	Type       string                      `yaml:"Type"`
	Properties LogsLogGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LogsLogGroupProperties struct {
	LogGroupName interface{} `yaml:"LogGroupName,omitempty"`
	RetentionInDays interface{} `yaml:"RetentionInDays,omitempty"`
}

func NewLogsLogGroup(properties LogsLogGroupProperties, deps ...interface{}) LogsLogGroup {
	return LogsLogGroup{
		Type:       "AWS::Logs::LogGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLogsLogGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource LogsLogGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsLogGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LogsLogGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LogsLogGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
