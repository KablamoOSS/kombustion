package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type SSMParameter struct {
	Type       string                      `yaml:"Type"`
	Properties SSMParameterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SSMParameterProperties struct {
	AllowedPattern interface{} `yaml:"AllowedPattern,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Type interface{} `yaml:"Type"`
	Value interface{} `yaml:"Value"`
}

func NewSSMParameter(properties SSMParameterProperties, deps ...interface{}) SSMParameter {
	return SSMParameter{
		Type:       "AWS::SSM::Parameter",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSSMParameter(name string, data string) (cf types.ValueMap, err error) {
	var resource SSMParameter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMParameter - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SSMParameter) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SSMParameterProperties) Validate() []error {
	errs := []error{}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
