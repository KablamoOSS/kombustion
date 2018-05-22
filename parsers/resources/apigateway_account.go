package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ApiGatewayAccount struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayAccountProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayAccountProperties struct {
	CloudWatchRoleArn interface{} `yaml:"CloudWatchRoleArn,omitempty"`
}

func NewApiGatewayAccount(properties ApiGatewayAccountProperties, deps ...interface{}) ApiGatewayAccount {
	return ApiGatewayAccount{
		Type:       "AWS::ApiGateway::Account",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayAccount(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayAccount
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayAccount - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayAccount) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayAccountProperties) Validate() []error {
	errs := []error{}
	return errs
}
