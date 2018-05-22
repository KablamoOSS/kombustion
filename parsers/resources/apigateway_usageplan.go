package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ApiGatewayUsagePlan struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayUsagePlanProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayUsagePlanProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	UsagePlanName interface{} `yaml:"UsagePlanName,omitempty"`
	Throttle *properties.UsagePlan_ThrottleSettings `yaml:"Throttle,omitempty"`
	Quota *properties.UsagePlan_QuotaSettings `yaml:"Quota,omitempty"`
	ApiStages interface{} `yaml:"ApiStages,omitempty"`
}

func NewApiGatewayUsagePlan(properties ApiGatewayUsagePlanProperties, deps ...interface{}) ApiGatewayUsagePlan {
	return ApiGatewayUsagePlan{
		Type:       "AWS::ApiGateway::UsagePlan",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayUsagePlan(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayUsagePlan
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayUsagePlan - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayUsagePlan) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayUsagePlanProperties) Validate() []error {
	errs := []error{}
	return errs
}
