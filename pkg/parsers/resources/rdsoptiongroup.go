package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// RDSOptionGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
type RDSOptionGroup struct {
	Type       string                   `yaml:"Type"`
	Properties RDSOptionGroupProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// RDSOptionGroup Properties
type RDSOptionGroupProperties struct {
	EngineName             interface{} `yaml:"EngineName"`
	MajorEngineVersion     interface{} `yaml:"MajorEngineVersion"`
	OptionGroupDescription interface{} `yaml:"OptionGroupDescription"`
	OptionConfigurations   interface{} `yaml:"OptionConfigurations"`
	Tags                   interface{} `yaml:"Tags,omitempty"`
}

// NewRDSOptionGroup constructor creates a new RDSOptionGroup
func NewRDSOptionGroup(properties RDSOptionGroupProperties, deps ...interface{}) RDSOptionGroup {
	return RDSOptionGroup{
		Type:       "AWS::RDS::OptionGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRDSOptionGroup parses RDSOptionGroup
func ParseRDSOptionGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
