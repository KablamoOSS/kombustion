package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EMRStep struct {
	Type       string                      `yaml:"Type"`
	Properties EMRStepProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EMRStepProperties struct {
	ActionOnFailure interface{} `yaml:"ActionOnFailure"`
	JobFlowId interface{} `yaml:"JobFlowId"`
	Name interface{} `yaml:"Name"`
	HadoopJarStep *properties.Step_HadoopJarStepConfig `yaml:"HadoopJarStep"`
}

func NewEMRStep(properties EMRStepProperties, deps ...interface{}) EMRStep {
	return EMRStep{
		Type:       "AWS::EMR::Step",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEMRStep(name string, data string) (cf types.ValueMap, err error) {
	var resource EMRStep
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRStep - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EMRStep) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EMRStepProperties) Validate() []error {
	errs := []error{}
	if resource.ActionOnFailure == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ActionOnFailure'"))
	}
	if resource.JobFlowId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'JobFlowId'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.HadoopJarStep == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HadoopJarStep'"))
	} else {
		errs = append(errs, resource.HadoopJarStep.Validate()...)
	}
	return errs
}
