package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayDocumentationVersion struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayDocumentationVersionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayDocumentationVersionProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	DocumentationVersion interface{} `yaml:"DocumentationVersion"`
	RestApiId interface{} `yaml:"RestApiId"`
}

func NewApiGatewayDocumentationVersion(properties ApiGatewayDocumentationVersionProperties, deps ...interface{}) ApiGatewayDocumentationVersion {
	return ApiGatewayDocumentationVersion{
		Type:       "AWS::ApiGateway::DocumentationVersion",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayDocumentationVersion(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayDocumentationVersion
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayDocumentationVersion - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayDocumentationVersion) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayDocumentationVersionProperties) Validate() []error {
	errs := []error{}
	if resource.DocumentationVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DocumentationVersion'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
