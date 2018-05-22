package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IAMUserToGroupAddition struct {
	Type       string                      `yaml:"Type"`
	Properties IAMUserToGroupAdditionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMUserToGroupAdditionProperties struct {
	GroupName interface{} `yaml:"GroupName"`
	Users interface{} `yaml:"Users"`
}

func NewIAMUserToGroupAddition(properties IAMUserToGroupAdditionProperties, deps ...interface{}) IAMUserToGroupAddition {
	return IAMUserToGroupAddition{
		Type:       "AWS::IAM::UserToGroupAddition",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMUserToGroupAddition(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMUserToGroupAddition
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMUserToGroupAddition - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMUserToGroupAddition) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMUserToGroupAdditionProperties) Validate() []error {
	errs := []error{}
	if resource.GroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupName'"))
	}
	if resource.Users == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Users'"))
	}
	return errs
}
