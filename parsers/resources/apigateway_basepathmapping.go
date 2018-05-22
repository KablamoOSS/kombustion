package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayBasePathMapping struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayBasePathMappingProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayBasePathMappingProperties struct {
	BasePath interface{} `yaml:"BasePath,omitempty"`
	DomainName interface{} `yaml:"DomainName"`
	RestApiId interface{} `yaml:"RestApiId,omitempty"`
	Stage interface{} `yaml:"Stage,omitempty"`
}

func NewApiGatewayBasePathMapping(properties ApiGatewayBasePathMappingProperties, deps ...interface{}) ApiGatewayBasePathMapping {
	return ApiGatewayBasePathMapping{
		Type:       "AWS::ApiGateway::BasePathMapping",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayBasePathMapping(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayBasePathMapping
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayBasePathMapping - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayBasePathMapping) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayBasePathMappingProperties) Validate() []error {
	errs := []error{}
	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	return errs
}
