package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EMRInstanceFleetConfig struct {
	Type       string                      `yaml:"Type"`
	Properties EMRInstanceFleetConfigProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EMRInstanceFleetConfigProperties struct {
	ClusterId interface{} `yaml:"ClusterId"`
	InstanceFleetType interface{} `yaml:"InstanceFleetType"`
	Name interface{} `yaml:"Name,omitempty"`
	TargetOnDemandCapacity interface{} `yaml:"TargetOnDemandCapacity,omitempty"`
	TargetSpotCapacity interface{} `yaml:"TargetSpotCapacity,omitempty"`
	InstanceTypeConfigs interface{} `yaml:"InstanceTypeConfigs,omitempty"`
	LaunchSpecifications *properties.InstanceFleetConfig_InstanceFleetProvisioningSpecifications `yaml:"LaunchSpecifications,omitempty"`
}

func NewEMRInstanceFleetConfig(properties EMRInstanceFleetConfigProperties, deps ...interface{}) EMRInstanceFleetConfig {
	return EMRInstanceFleetConfig{
		Type:       "AWS::EMR::InstanceFleetConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEMRInstanceFleetConfig(name string, data string) (cf types.ValueMap, err error) {
	var resource EMRInstanceFleetConfig
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRInstanceFleetConfig - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EMRInstanceFleetConfig) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EMRInstanceFleetConfigProperties) Validate() []error {
	errs := []error{}
	if resource.ClusterId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ClusterId'"))
	}
	if resource.InstanceFleetType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceFleetType'"))
	}
	return errs
}
