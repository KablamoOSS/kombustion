package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayResource struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayResourceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayResourceProperties struct {
	ParentId interface{} `yaml:"ParentId"`
	PathPart interface{} `yaml:"PathPart"`
	RestApiId interface{} `yaml:"RestApiId"`
}

func NewApiGatewayResource(properties ApiGatewayResourceProperties, deps ...interface{}) ApiGatewayResource {
	return ApiGatewayResource{
		Type:       "AWS::ApiGateway::Resource",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayResource(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayResource
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayResource - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayResource) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayResourceProperties) Validate() []error {
	errs := []error{}
	if resource.ParentId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParentId'"))
	}
	if resource.PathPart == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PathPart'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
