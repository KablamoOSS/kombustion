package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ApiGatewayDeployment struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayDeploymentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayDeploymentProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	RestApiId interface{} `yaml:"RestApiId"`
	StageName interface{} `yaml:"StageName,omitempty"`
	StageDescription *properties.Deployment_StageDescription `yaml:"StageDescription,omitempty"`
}

func NewApiGatewayDeployment(properties ApiGatewayDeploymentProperties, deps ...interface{}) ApiGatewayDeployment {
	return ApiGatewayDeployment{
		Type:       "AWS::ApiGateway::Deployment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayDeployment(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayDeployment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayDeployment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayDeployment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayDeploymentProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
