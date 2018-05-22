package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayRequestValidator struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayRequestValidatorProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayRequestValidatorProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	RestApiId interface{} `yaml:"RestApiId"`
	ValidateRequestBody interface{} `yaml:"ValidateRequestBody,omitempty"`
	ValidateRequestParameters interface{} `yaml:"ValidateRequestParameters,omitempty"`
}

func NewApiGatewayRequestValidator(properties ApiGatewayRequestValidatorProperties, deps ...interface{}) ApiGatewayRequestValidator {
	return ApiGatewayRequestValidator{
		Type:       "AWS::ApiGateway::RequestValidator",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayRequestValidator(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayRequestValidator
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayRequestValidator - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayRequestValidator) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayRequestValidatorProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
