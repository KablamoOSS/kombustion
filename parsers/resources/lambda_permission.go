package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type LambdaPermission struct {
	Type       string                      `yaml:"Type"`
	Properties LambdaPermissionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LambdaPermissionProperties struct {
	Action interface{} `yaml:"Action"`
	EventSourceToken interface{} `yaml:"EventSourceToken,omitempty"`
	FunctionName interface{} `yaml:"FunctionName"`
	Principal interface{} `yaml:"Principal"`
	SourceAccount interface{} `yaml:"SourceAccount,omitempty"`
	SourceArn interface{} `yaml:"SourceArn,omitempty"`
}

func NewLambdaPermission(properties LambdaPermissionProperties, deps ...interface{}) LambdaPermission {
	return LambdaPermission{
		Type:       "AWS::Lambda::Permission",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLambdaPermission(name string, data string) (cf types.ValueMap, err error) {
	var resource LambdaPermission
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaPermission - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LambdaPermission) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LambdaPermissionProperties) Validate() []error {
	errs := []error{}
	if resource.Action == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Action'"))
	}
	if resource.FunctionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	if resource.Principal == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Principal'"))
	}
	return errs
}
