package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CloudFormationWaitCondition struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFormationWaitConditionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFormationWaitConditionProperties struct {
	Count interface{} `yaml:"Count,omitempty"`
	Handle interface{} `yaml:"Handle"`
	Timeout interface{} `yaml:"Timeout"`
}

func NewCloudFormationWaitCondition(properties CloudFormationWaitConditionProperties, deps ...interface{}) CloudFormationWaitCondition {
	return CloudFormationWaitCondition{
		Type:       "AWS::CloudFormation::WaitCondition",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFormationWaitCondition(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFormationWaitCondition
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFormationWaitCondition - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFormationWaitCondition) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFormationWaitConditionProperties) Validate() []error {
	errs := []error{}
	if resource.Handle == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Handle'"))
	}
	if resource.Timeout == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Timeout'"))
	}
	return errs
}
