package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayModel struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayModelProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayModelProperties struct {
	ContentType interface{} `yaml:"ContentType,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	RestApiId interface{} `yaml:"RestApiId"`
	Schema interface{} `yaml:"Schema,omitempty"`
}

func NewApiGatewayModel(properties ApiGatewayModelProperties, deps ...interface{}) ApiGatewayModel {
	return ApiGatewayModel{
		Type:       "AWS::ApiGateway::Model",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayModel(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayModel
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayModel - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayModel) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayModelProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
