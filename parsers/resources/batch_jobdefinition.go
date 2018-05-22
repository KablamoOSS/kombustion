package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type BatchJobDefinition struct {
	Type       string                      `yaml:"Type"`
	Properties BatchJobDefinitionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type BatchJobDefinitionProperties struct {
	JobDefinitionName interface{} `yaml:"JobDefinitionName,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	Type interface{} `yaml:"Type"`
	RetryStrategy *properties.JobDefinition_RetryStrategy `yaml:"RetryStrategy,omitempty"`
	ContainerProperties *properties.JobDefinition_ContainerProperties `yaml:"ContainerProperties"`
}

func NewBatchJobDefinition(properties BatchJobDefinitionProperties, deps ...interface{}) BatchJobDefinition {
	return BatchJobDefinition{
		Type:       "AWS::Batch::JobDefinition",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseBatchJobDefinition(name string, data string) (cf types.ValueMap, err error) {
	var resource BatchJobDefinition
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: BatchJobDefinition - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource BatchJobDefinition) Validate() []error {
	return resource.Properties.Validate()
}

func (resource BatchJobDefinitionProperties) Validate() []error {
	errs := []error{}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.ContainerProperties == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ContainerProperties'"))
	} else {
		errs = append(errs, resource.ContainerProperties.Validate()...)
	}
	return errs
}
