package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2FlowLog struct {
	Type       string                      `yaml:"Type"`
	Properties EC2FlowLogProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2FlowLogProperties struct {
	DeliverLogsPermissionArn interface{} `yaml:"DeliverLogsPermissionArn"`
	LogGroupName interface{} `yaml:"LogGroupName"`
	ResourceId interface{} `yaml:"ResourceId"`
	ResourceType interface{} `yaml:"ResourceType"`
	TrafficType interface{} `yaml:"TrafficType"`
}

func NewEC2FlowLog(properties EC2FlowLogProperties, deps ...interface{}) EC2FlowLog {
	return EC2FlowLog{
		Type:       "AWS::EC2::FlowLog",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2FlowLog(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2FlowLog
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2FlowLog - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2FlowLog) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2FlowLogProperties) Validate() []error {
	errs := []error{}
	if resource.DeliverLogsPermissionArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeliverLogsPermissionArn'"))
	}
	if resource.LogGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LogGroupName'"))
	}
	if resource.ResourceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceId'"))
	}
	if resource.ResourceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceType'"))
	}
	if resource.TrafficType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TrafficType'"))
	}
	return errs
}
