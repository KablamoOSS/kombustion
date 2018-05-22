package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CodeDeployDeploymentConfig struct {
	Type       string                      `yaml:"Type"`
	Properties CodeDeployDeploymentConfigProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodeDeployDeploymentConfigProperties struct {
	DeploymentConfigName interface{} `yaml:"DeploymentConfigName,omitempty"`
	MinimumHealthyHosts *properties.DeploymentConfig_MinimumHealthyHosts `yaml:"MinimumHealthyHosts,omitempty"`
}

func NewCodeDeployDeploymentConfig(properties CodeDeployDeploymentConfigProperties, deps ...interface{}) CodeDeployDeploymentConfig {
	return CodeDeployDeploymentConfig{
		Type:       "AWS::CodeDeploy::DeploymentConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodeDeployDeploymentConfig(name string, data string) (cf types.ValueMap, err error) {
	var resource CodeDeployDeploymentConfig
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeDeployDeploymentConfig - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodeDeployDeploymentConfig) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodeDeployDeploymentConfigProperties) Validate() []error {
	errs := []error{}
	return errs
}
