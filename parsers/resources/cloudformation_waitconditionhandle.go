package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type CloudFormationWaitConditionHandle struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFormationWaitConditionHandleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFormationWaitConditionHandleProperties struct {
}

func NewCloudFormationWaitConditionHandle(properties CloudFormationWaitConditionHandleProperties, deps ...interface{}) CloudFormationWaitConditionHandle {
	return CloudFormationWaitConditionHandle{
		Type:       "AWS::CloudFormation::WaitConditionHandle",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFormationWaitConditionHandle(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFormationWaitConditionHandle
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFormationWaitConditionHandle - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFormationWaitConditionHandle) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFormationWaitConditionHandleProperties) Validate() []error {
	errs := []error{}
	return errs
}
