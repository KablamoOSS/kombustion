package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayUsagePlanKey struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayUsagePlanKeyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayUsagePlanKeyProperties struct {
	KeyId interface{} `yaml:"KeyId"`
	KeyType interface{} `yaml:"KeyType"`
	UsagePlanId interface{} `yaml:"UsagePlanId"`
}

func NewApiGatewayUsagePlanKey(properties ApiGatewayUsagePlanKeyProperties, deps ...interface{}) ApiGatewayUsagePlanKey {
	return ApiGatewayUsagePlanKey{
		Type:       "AWS::ApiGateway::UsagePlanKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayUsagePlanKey(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayUsagePlanKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayUsagePlanKey - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayUsagePlanKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayUsagePlanKeyProperties) Validate() []error {
	errs := []error{}
	if resource.KeyId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyId'"))
	}
	if resource.KeyType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyType'"))
	}
	if resource.UsagePlanId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UsagePlanId'"))
	}
	return errs
}
