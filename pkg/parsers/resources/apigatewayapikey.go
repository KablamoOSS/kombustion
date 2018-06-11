package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayApiKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
type ApiGatewayApiKey struct {
	Type       string                     `yaml:"Type"`
	Properties ApiGatewayApiKeyProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// ApiGatewayApiKey Properties
type ApiGatewayApiKeyProperties struct {
	CustomerId         interface{} `yaml:"CustomerId,omitempty"`
	Description        interface{} `yaml:"Description,omitempty"`
	Enabled            interface{} `yaml:"Enabled,omitempty"`
	GenerateDistinctId interface{} `yaml:"GenerateDistinctId,omitempty"`
	Name               interface{} `yaml:"Name,omitempty"`
	StageKeys          interface{} `yaml:"StageKeys,omitempty"`
}

// NewApiGatewayApiKey constructor creates a new ApiGatewayApiKey
func NewApiGatewayApiKey(properties ApiGatewayApiKeyProperties, deps ...interface{}) ApiGatewayApiKey {
	return ApiGatewayApiKey{
		Type:       "AWS::ApiGateway::ApiKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayApiKey parses ApiGatewayApiKey
func ParseApiGatewayApiKey(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayApiKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayApiKey - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ApiGatewayApiKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayApiKeyProperties) Validate() []error {
	errs := []error{}
	return errs
}
