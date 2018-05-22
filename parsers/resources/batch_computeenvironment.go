package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type BatchComputeEnvironment struct {
	Type       string                      `yaml:"Type"`
	Properties BatchComputeEnvironmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type BatchComputeEnvironmentProperties struct {
	ComputeEnvironmentName interface{} `yaml:"ComputeEnvironmentName,omitempty"`
	ServiceRole interface{} `yaml:"ServiceRole"`
	State interface{} `yaml:"State,omitempty"`
	Type interface{} `yaml:"Type"`
	ComputeResources *properties.ComputeEnvironment_ComputeResources `yaml:"ComputeResources"`
}

func NewBatchComputeEnvironment(properties BatchComputeEnvironmentProperties, deps ...interface{}) BatchComputeEnvironment {
	return BatchComputeEnvironment{
		Type:       "AWS::Batch::ComputeEnvironment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseBatchComputeEnvironment(name string, data string) (cf types.ValueMap, err error) {
	var resource BatchComputeEnvironment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: BatchComputeEnvironment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource BatchComputeEnvironment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource BatchComputeEnvironmentProperties) Validate() []error {
	errs := []error{}
	if resource.ServiceRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRole'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.ComputeResources == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComputeResources'"))
	} else {
		errs = append(errs, resource.ComputeResources.Validate()...)
	}
	return errs
}
