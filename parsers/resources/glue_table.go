package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueTable struct {
	Type       string                      `yaml:"Type"`
	Properties GlueTableProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueTableProperties struct {
	CatalogId interface{} `yaml:"CatalogId"`
	DatabaseName interface{} `yaml:"DatabaseName"`
	TableInput *properties.Table_TableInput `yaml:"TableInput"`
}

func NewGlueTable(properties GlueTableProperties, deps ...interface{}) GlueTable {
	return GlueTable{
		Type:       "AWS::Glue::Table",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueTable(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueTable
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueTable - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueTable) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueTableProperties) Validate() []error {
	errs := []error{}
	if resource.CatalogId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.DatabaseName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DatabaseName'"))
	}
	if resource.TableInput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableInput'"))
	} else {
		errs = append(errs, resource.TableInput.Validate()...)
	}
	return errs
}
