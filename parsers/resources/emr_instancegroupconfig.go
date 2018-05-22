package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EMRInstanceGroupConfig struct {
	Type       string                      `yaml:"Type"`
	Properties EMRInstanceGroupConfigProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EMRInstanceGroupConfigProperties struct {
	BidPrice interface{} `yaml:"BidPrice,omitempty"`
	InstanceCount interface{} `yaml:"InstanceCount"`
	InstanceRole interface{} `yaml:"InstanceRole"`
	InstanceType interface{} `yaml:"InstanceType"`
	JobFlowId interface{} `yaml:"JobFlowId"`
	Market interface{} `yaml:"Market,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Configurations interface{} `yaml:"Configurations,omitempty"`
	EbsConfiguration *properties.InstanceGroupConfig_EbsConfiguration `yaml:"EbsConfiguration,omitempty"`
	AutoScalingPolicy *properties.InstanceGroupConfig_AutoScalingPolicy `yaml:"AutoScalingPolicy,omitempty"`
}

func NewEMRInstanceGroupConfig(properties EMRInstanceGroupConfigProperties, deps ...interface{}) EMRInstanceGroupConfig {
	return EMRInstanceGroupConfig{
		Type:       "AWS::EMR::InstanceGroupConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEMRInstanceGroupConfig(name string, data string) (cf types.ValueMap, err error) {
	var resource EMRInstanceGroupConfig
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRInstanceGroupConfig - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EMRInstanceGroupConfig) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EMRInstanceGroupConfigProperties) Validate() []error {
	errs := []error{}
	if resource.InstanceCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceCount'"))
	}
	if resource.InstanceRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceRole'"))
	}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	if resource.JobFlowId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'JobFlowId'"))
	}
	return errs
}
