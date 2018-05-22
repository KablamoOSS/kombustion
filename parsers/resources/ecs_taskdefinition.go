package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ECSTaskDefinition struct {
	Type       string                      `yaml:"Type"`
	Properties ECSTaskDefinitionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ECSTaskDefinitionProperties struct {
	Cpu interface{} `yaml:"Cpu,omitempty"`
	ExecutionRoleArn interface{} `yaml:"ExecutionRoleArn,omitempty"`
	Family interface{} `yaml:"Family,omitempty"`
	Memory interface{} `yaml:"Memory,omitempty"`
	NetworkMode interface{} `yaml:"NetworkMode,omitempty"`
	TaskRoleArn interface{} `yaml:"TaskRoleArn,omitempty"`
	ContainerDefinitions interface{} `yaml:"ContainerDefinitions,omitempty"`
	PlacementConstraints interface{} `yaml:"PlacementConstraints,omitempty"`
	RequiresCompatibilities interface{} `yaml:"RequiresCompatibilities,omitempty"`
	Volumes interface{} `yaml:"Volumes,omitempty"`
}

func NewECSTaskDefinition(properties ECSTaskDefinitionProperties, deps ...interface{}) ECSTaskDefinition {
	return ECSTaskDefinition{
		Type:       "AWS::ECS::TaskDefinition",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseECSTaskDefinition(name string, data string) (cf types.ValueMap, err error) {
	var resource ECSTaskDefinition
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ECSTaskDefinition - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ECSTaskDefinition) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ECSTaskDefinitionProperties) Validate() []error {
	errs := []error{}
	return errs
}
