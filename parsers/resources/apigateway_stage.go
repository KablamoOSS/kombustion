package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ApiGatewayStage struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayStageProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayStageProperties struct {
	CacheClusterEnabled interface{} `yaml:"CacheClusterEnabled,omitempty"`
	CacheClusterSize interface{} `yaml:"CacheClusterSize,omitempty"`
	ClientCertificateId interface{} `yaml:"ClientCertificateId,omitempty"`
	DeploymentId interface{} `yaml:"DeploymentId,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	DocumentationVersion interface{} `yaml:"DocumentationVersion,omitempty"`
	RestApiId interface{} `yaml:"RestApiId"`
	StageName interface{} `yaml:"StageName,omitempty"`
	Variables interface{} `yaml:"Variables,omitempty"`
	MethodSettings interface{} `yaml:"MethodSettings,omitempty"`
}

func NewApiGatewayStage(properties ApiGatewayStageProperties, deps ...interface{}) ApiGatewayStage {
	return ApiGatewayStage{
		Type:       "AWS::ApiGateway::Stage",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayStage(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayStage
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayStage - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayStage) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayStageProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
