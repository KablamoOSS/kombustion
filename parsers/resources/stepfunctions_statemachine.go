package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type StepFunctionsStateMachine struct {
	Type       string                      `yaml:"Type"`
	Properties StepFunctionsStateMachineProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type StepFunctionsStateMachineProperties struct {
	DefinitionString interface{} `yaml:"DefinitionString"`
	RoleArn interface{} `yaml:"RoleArn"`
	StateMachineName interface{} `yaml:"StateMachineName,omitempty"`
}

func NewStepFunctionsStateMachine(properties StepFunctionsStateMachineProperties, deps ...interface{}) StepFunctionsStateMachine {
	return StepFunctionsStateMachine{
		Type:       "AWS::StepFunctions::StateMachine",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseStepFunctionsStateMachine(name string, data string) (cf types.ValueMap, err error) {
	var resource StepFunctionsStateMachine
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: StepFunctionsStateMachine - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource StepFunctionsStateMachine) Validate() []error {
	return resource.Properties.Validate()
}

func (resource StepFunctionsStateMachineProperties) Validate() []error {
	errs := []error{}
	if resource.DefinitionString == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DefinitionString'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
