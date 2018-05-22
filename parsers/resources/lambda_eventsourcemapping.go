package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type LambdaEventSourceMapping struct {
	Type       string                      `yaml:"Type"`
	Properties LambdaEventSourceMappingProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LambdaEventSourceMappingProperties struct {
	BatchSize interface{} `yaml:"BatchSize,omitempty"`
	Enabled interface{} `yaml:"Enabled,omitempty"`
	EventSourceArn interface{} `yaml:"EventSourceArn"`
	FunctionName interface{} `yaml:"FunctionName"`
	StartingPosition interface{} `yaml:"StartingPosition"`
}

func NewLambdaEventSourceMapping(properties LambdaEventSourceMappingProperties, deps ...interface{}) LambdaEventSourceMapping {
	return LambdaEventSourceMapping{
		Type:       "AWS::Lambda::EventSourceMapping",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLambdaEventSourceMapping(name string, data string) (cf types.ValueMap, err error) {
	var resource LambdaEventSourceMapping
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaEventSourceMapping - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LambdaEventSourceMapping) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LambdaEventSourceMappingProperties) Validate() []error {
	errs := []error{}
	if resource.EventSourceArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EventSourceArn'"))
	}
	if resource.FunctionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	if resource.StartingPosition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StartingPosition'"))
	}
	return errs
}
