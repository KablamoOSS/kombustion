package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ECSCluster struct {
	Type       string                      `yaml:"Type"`
	Properties ECSClusterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ECSClusterProperties struct {
	ClusterName interface{} `yaml:"ClusterName,omitempty"`
}

func NewECSCluster(properties ECSClusterProperties, deps ...interface{}) ECSCluster {
	return ECSCluster{
		Type:       "AWS::ECS::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseECSCluster(name string, data string) (cf types.ValueMap, err error) {
	var resource ECSCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ECSCluster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ECSCluster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ECSClusterProperties) Validate() []error {
	errs := []error{}
	return errs
}
