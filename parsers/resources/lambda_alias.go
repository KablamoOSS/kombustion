package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type LambdaAlias struct {
	Type       string                      `yaml:"Type"`
	Properties LambdaAliasProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LambdaAliasProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	FunctionName interface{} `yaml:"FunctionName"`
	FunctionVersion interface{} `yaml:"FunctionVersion"`
	Name interface{} `yaml:"Name"`
	RoutingConfig *properties.Alias_AliasRoutingConfiguration `yaml:"RoutingConfig,omitempty"`
}

func NewLambdaAlias(properties LambdaAliasProperties, deps ...interface{}) LambdaAlias {
	return LambdaAlias{
		Type:       "AWS::Lambda::Alias",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLambdaAlias(name string, data string) (cf types.ValueMap, err error) {
	var resource LambdaAlias
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaAlias - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LambdaAlias) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LambdaAliasProperties) Validate() []error {
	errs := []error{}
	if resource.FunctionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	if resource.FunctionVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionVersion'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
