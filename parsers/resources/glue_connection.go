package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueConnection struct {
	Type       string                      `yaml:"Type"`
	Properties GlueConnectionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueConnectionProperties struct {
	CatalogId interface{} `yaml:"CatalogId"`
	ConnectionInput *properties.Connection_ConnectionInput `yaml:"ConnectionInput"`
}

func NewGlueConnection(properties GlueConnectionProperties, deps ...interface{}) GlueConnection {
	return GlueConnection{
		Type:       "AWS::Glue::Connection",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueConnection(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueConnection
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueConnection - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueConnection) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueConnectionProperties) Validate() []error {
	errs := []error{}
	if resource.CatalogId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.ConnectionInput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ConnectionInput'"))
	} else {
		errs = append(errs, resource.ConnectionInput.Validate()...)
	}
	return errs
}
