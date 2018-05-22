package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type LambdaVersion struct {
	Type       string                      `yaml:"Type"`
	Properties LambdaVersionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LambdaVersionProperties struct {
	CodeSha256 interface{} `yaml:"CodeSha256,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	FunctionName interface{} `yaml:"FunctionName"`
}

func NewLambdaVersion(properties LambdaVersionProperties, deps ...interface{}) LambdaVersion {
	return LambdaVersion{
		Type:       "AWS::Lambda::Version",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLambdaVersion(name string, data string) (cf types.ValueMap, err error) {
	var resource LambdaVersion
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaVersion - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LambdaVersion) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LambdaVersionProperties) Validate() []error {
	errs := []error{}
	if resource.FunctionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	return errs
}
