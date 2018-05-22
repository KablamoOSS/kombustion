package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueDatabase struct {
	Type       string                      `yaml:"Type"`
	Properties GlueDatabaseProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueDatabaseProperties struct {
	CatalogId interface{} `yaml:"CatalogId"`
	DatabaseInput *properties.Database_DatabaseInput `yaml:"DatabaseInput"`
}

func NewGlueDatabase(properties GlueDatabaseProperties, deps ...interface{}) GlueDatabase {
	return GlueDatabase{
		Type:       "AWS::Glue::Database",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueDatabase(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueDatabase
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueDatabase - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueDatabase) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueDatabaseProperties) Validate() []error {
	errs := []error{}
	if resource.CatalogId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.DatabaseInput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DatabaseInput'"))
	} else {
		errs = append(errs, resource.DatabaseInput.Validate()...)
	}
	return errs
}
