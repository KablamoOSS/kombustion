package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type OpsWorksLayer struct {
	Type       string                      `yaml:"Type"`
	Properties OpsWorksLayerProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type OpsWorksLayerProperties struct {
	AutoAssignElasticIps interface{} `yaml:"AutoAssignElasticIps"`
	AutoAssignPublicIps interface{} `yaml:"AutoAssignPublicIps"`
	CustomInstanceProfileArn interface{} `yaml:"CustomInstanceProfileArn,omitempty"`
	CustomJson interface{} `yaml:"CustomJson,omitempty"`
	EnableAutoHealing interface{} `yaml:"EnableAutoHealing"`
	InstallUpdatesOnBoot interface{} `yaml:"InstallUpdatesOnBoot,omitempty"`
	Name interface{} `yaml:"Name"`
	Shortname interface{} `yaml:"Shortname"`
	StackId interface{} `yaml:"StackId"`
	Type interface{} `yaml:"Type"`
	UseEbsOptimizedInstances interface{} `yaml:"UseEbsOptimizedInstances,omitempty"`
	CustomRecipes *properties.Layer_Recipes `yaml:"CustomRecipes,omitempty"`
	Attributes interface{} `yaml:"Attributes,omitempty"`
	LoadBasedAutoScaling *properties.Layer_LoadBasedAutoScaling `yaml:"LoadBasedAutoScaling,omitempty"`
	CustomSecurityGroupIds interface{} `yaml:"CustomSecurityGroupIds,omitempty"`
	Packages interface{} `yaml:"Packages,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VolumeConfigurations interface{} `yaml:"VolumeConfigurations,omitempty"`
	LifecycleEventConfiguration *properties.Layer_LifecycleEventConfiguration `yaml:"LifecycleEventConfiguration,omitempty"`
}

func NewOpsWorksLayer(properties OpsWorksLayerProperties, deps ...interface{}) OpsWorksLayer {
	return OpsWorksLayer{
		Type:       "AWS::OpsWorks::Layer",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksLayer(name string, data string) (cf types.ValueMap, err error) {
	var resource OpsWorksLayer
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksLayer - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource OpsWorksLayer) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksLayerProperties) Validate() []error {
	errs := []error{}
	if resource.AutoAssignElasticIps == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoAssignElasticIps'"))
	}
	if resource.AutoAssignPublicIps == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoAssignPublicIps'"))
	}
	if resource.EnableAutoHealing == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EnableAutoHealing'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Shortname == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Shortname'"))
	}
	if resource.StackId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StackId'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
