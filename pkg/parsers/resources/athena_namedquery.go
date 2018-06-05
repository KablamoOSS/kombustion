package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type AthenaNamedQuery struct {
	Type       string                     `yaml:"Type"`
	Properties AthenaNamedQueryProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

type AthenaNamedQueryProperties struct {
	Database    interface{} `yaml:"Database"`
	Description interface{} `yaml:"Description,omitempty"`
	Name        interface{} `yaml:"Name,omitempty"`
	QueryString interface{} `yaml:"QueryString"`
}

func NewAthenaNamedQuery(properties AthenaNamedQueryProperties, deps ...interface{}) AthenaNamedQuery {
	return AthenaNamedQuery{
		Type:       "AWS::Athena::NamedQuery",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAthenaNamedQuery(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource AthenaNamedQuery
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AthenaNamedQuery - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource AthenaNamedQuery) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AthenaNamedQueryProperties) Validate() []error {
	errs := []error{}
	if resource.Database == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Database'"))
	}
	if resource.QueryString == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'QueryString'"))
	}
	return errs
}