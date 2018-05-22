package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ApiGatewayDocumentationPart struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayDocumentationPartProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayDocumentationPartProperties struct {
	Properties interface{} `yaml:"Properties"`
	RestApiId interface{} `yaml:"RestApiId"`
	Location *properties.DocumentationPart_Location `yaml:"Location"`
}

func NewApiGatewayDocumentationPart(properties ApiGatewayDocumentationPartProperties, deps ...interface{}) ApiGatewayDocumentationPart {
	return ApiGatewayDocumentationPart{
		Type:       "AWS::ApiGateway::DocumentationPart",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayDocumentationPart(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayDocumentationPart
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayDocumentationPart - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayDocumentationPart) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayDocumentationPartProperties) Validate() []error {
	errs := []error{}
	if resource.Properties == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Properties'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	if resource.Location == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Location'"))
	} else {
		errs = append(errs, resource.Location.Validate()...)
	}
	return errs
}
