package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type StepFunctionsActivity struct {
	Type       string                      `yaml:"Type"`
	Properties StepFunctionsActivityProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type StepFunctionsActivityProperties struct {
	Name interface{} `yaml:"Name"`
}

func NewStepFunctionsActivity(properties StepFunctionsActivityProperties, deps ...interface{}) StepFunctionsActivity {
	return StepFunctionsActivity{
		Type:       "AWS::StepFunctions::Activity",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseStepFunctionsActivity(name string, data string) (cf types.ValueMap, err error) {
	var resource StepFunctionsActivity
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: StepFunctionsActivity - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource StepFunctionsActivity) Validate() []error {
	return resource.Properties.Validate()
}

func (resource StepFunctionsActivityProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
