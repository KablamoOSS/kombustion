package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GluePartition struct {
	Type       string                      `yaml:"Type"`
	Properties GluePartitionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GluePartitionProperties struct {
	CatalogId interface{} `yaml:"CatalogId"`
	DatabaseName interface{} `yaml:"DatabaseName"`
	TableName interface{} `yaml:"TableName"`
	PartitionInput *properties.Partition_PartitionInput `yaml:"PartitionInput"`
}

func NewGluePartition(properties GluePartitionProperties, deps ...interface{}) GluePartition {
	return GluePartition{
		Type:       "AWS::Glue::Partition",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGluePartition(name string, data string) (cf types.ValueMap, err error) {
	var resource GluePartition
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GluePartition - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GluePartition) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GluePartitionProperties) Validate() []error {
	errs := []error{}
	if resource.CatalogId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.DatabaseName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DatabaseName'"))
	}
	if resource.TableName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableName'"))
	}
	if resource.PartitionInput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PartitionInput'"))
	} else {
		errs = append(errs, resource.PartitionInput.Validate()...)
	}
	return errs
}
