package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CodeDeployDeploymentGroup struct {
	Type       string                      `yaml:"Type"`
	Properties CodeDeployDeploymentGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodeDeployDeploymentGroupProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName"`
	DeploymentConfigName interface{} `yaml:"DeploymentConfigName,omitempty"`
	DeploymentGroupName interface{} `yaml:"DeploymentGroupName,omitempty"`
	ServiceRoleArn interface{} `yaml:"ServiceRoleArn"`
	LoadBalancerInfo *properties.DeploymentGroup_LoadBalancerInfo `yaml:"LoadBalancerInfo,omitempty"`
	OnPremisesInstanceTagFilters interface{} `yaml:"OnPremisesInstanceTagFilters,omitempty"`
	AutoScalingGroups interface{} `yaml:"AutoScalingGroups,omitempty"`
	Ec2TagFilters interface{} `yaml:"Ec2TagFilters,omitempty"`
	TriggerConfigurations interface{} `yaml:"TriggerConfigurations,omitempty"`
	DeploymentStyle *properties.DeploymentGroup_DeploymentStyle `yaml:"DeploymentStyle,omitempty"`
	Deployment *properties.DeploymentGroup_Deployment `yaml:"Deployment,omitempty"`
	AutoRollbackConfiguration *properties.DeploymentGroup_AutoRollbackConfiguration `yaml:"AutoRollbackConfiguration,omitempty"`
	AlarmConfiguration *properties.DeploymentGroup_AlarmConfiguration `yaml:"AlarmConfiguration,omitempty"`
}

func NewCodeDeployDeploymentGroup(properties CodeDeployDeploymentGroupProperties, deps ...interface{}) CodeDeployDeploymentGroup {
	return CodeDeployDeploymentGroup{
		Type:       "AWS::CodeDeploy::DeploymentGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodeDeployDeploymentGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource CodeDeployDeploymentGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeDeployDeploymentGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodeDeployDeploymentGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodeDeployDeploymentGroupProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.ServiceRoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRoleArn'"))
	}
	return errs
}
