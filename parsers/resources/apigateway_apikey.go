package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ApiGatewayApiKey struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayApiKeyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayApiKeyProperties struct {
	CustomerId interface{} `yaml:"CustomerId,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	Enabled interface{} `yaml:"Enabled,omitempty"`
	GenerateDistinctId interface{} `yaml:"GenerateDistinctId,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	StageKeys interface{} `yaml:"StageKeys,omitempty"`
}

func NewApiGatewayApiKey(properties ApiGatewayApiKeyProperties, deps ...interface{}) ApiGatewayApiKey {
	return ApiGatewayApiKey{
		Type:       "AWS::ApiGateway::ApiKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayApiKey(name string, data string) (cf types.ValueMap, err error) {
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
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayApiKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayApiKeyProperties) Validate() []error {
	errs := []error{}
	return errs
}
