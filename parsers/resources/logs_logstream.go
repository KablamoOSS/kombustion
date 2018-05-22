package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type LogsLogStream struct {
	Type       string                      `yaml:"Type"`
	Properties LogsLogStreamProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LogsLogStreamProperties struct {
	LogGroupName interface{} `yaml:"LogGroupName"`
	LogStreamName interface{} `yaml:"LogStreamName,omitempty"`
}

func NewLogsLogStream(properties LogsLogStreamProperties, deps ...interface{}) LogsLogStream {
	return LogsLogStream{
		Type:       "AWS::Logs::LogStream",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLogsLogStream(name string, data string) (cf types.ValueMap, err error) {
	var resource LogsLogStream
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsLogStream - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LogsLogStream) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LogsLogStreamProperties) Validate() []error {
	errs := []error{}
	if resource.LogGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LogGroupName'"))
	}
	return errs
}
