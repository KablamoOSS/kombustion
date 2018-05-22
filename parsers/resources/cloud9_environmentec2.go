package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type Cloud9EnvironmentEC2 struct {
	Type       string                      `yaml:"Type"`
	Properties Cloud9EnvironmentEC2Properties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type Cloud9EnvironmentEC2Properties struct {
	AutomaticStopTimeMinutes interface{} `yaml:"AutomaticStopTimeMinutes,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	InstanceType interface{} `yaml:"InstanceType"`
	Name interface{} `yaml:"Name,omitempty"`
	OwnerArn interface{} `yaml:"OwnerArn,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	Repositories interface{} `yaml:"Repositories,omitempty"`
}

func NewCloud9EnvironmentEC2(properties Cloud9EnvironmentEC2Properties, deps ...interface{}) Cloud9EnvironmentEC2 {
	return Cloud9EnvironmentEC2{
		Type:       "AWS::Cloud9::EnvironmentEC2",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloud9EnvironmentEC2(name string, data string) (cf types.ValueMap, err error) {
	var resource Cloud9EnvironmentEC2
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Cloud9EnvironmentEC2 - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource Cloud9EnvironmentEC2) Validate() []error {
	return resource.Properties.Validate()
}

func (resource Cloud9EnvironmentEC2Properties) Validate() []error {
	errs := []error{}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
