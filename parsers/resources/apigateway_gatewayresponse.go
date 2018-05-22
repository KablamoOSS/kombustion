package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayGatewayResponse struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayGatewayResponseProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayGatewayResponseProperties struct {
	ResponseType interface{} `yaml:"ResponseType"`
	RestApiId interface{} `yaml:"RestApiId"`
	StatusCode interface{} `yaml:"StatusCode,omitempty"`
	ResponseParameters interface{} `yaml:"ResponseParameters,omitempty"`
	ResponseTemplates interface{} `yaml:"ResponseTemplates,omitempty"`
}

func NewApiGatewayGatewayResponse(properties ApiGatewayGatewayResponseProperties, deps ...interface{}) ApiGatewayGatewayResponse {
	return ApiGatewayGatewayResponse{
		Type:       "AWS::ApiGateway::GatewayResponse",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayGatewayResponse(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayGatewayResponse
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayGatewayResponse - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayGatewayResponse) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayGatewayResponseProperties) Validate() []error {
	errs := []error{}
	if resource.ResponseType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResponseType'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
