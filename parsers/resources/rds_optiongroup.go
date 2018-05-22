package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSOptionGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RDSOptionGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSOptionGroupProperties struct {
	EngineName interface{} `yaml:"EngineName"`
	MajorEngineVersion interface{} `yaml:"MajorEngineVersion"`
	OptionGroupDescription interface{} `yaml:"OptionGroupDescription"`
	OptionConfigurations interface{} `yaml:"OptionConfigurations"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRDSOptionGroup(properties RDSOptionGroupProperties, deps ...interface{}) RDSOptionGroup {
	return RDSOptionGroup{
		Type:       "AWS::RDS::OptionGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSOptionGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSOptionGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSOptionGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSOptionGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSOptionGroupProperties) Validate() []error {
	errs := []error{}
	if resource.EngineName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EngineName'"))
	}
	if resource.MajorEngineVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MajorEngineVersion'"))
	}
	if resource.OptionGroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OptionGroupDescription'"))
	}
	if resource.OptionConfigurations == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OptionConfigurations'"))
	}
	return errs
}
