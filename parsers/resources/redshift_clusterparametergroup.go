package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RedshiftClusterParameterGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RedshiftClusterParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RedshiftClusterParameterGroupProperties struct {
	Description interface{} `yaml:"Description"`
	ParameterGroupFamily interface{} `yaml:"ParameterGroupFamily"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRedshiftClusterParameterGroup(properties RedshiftClusterParameterGroupProperties, deps ...interface{}) RedshiftClusterParameterGroup {
	return RedshiftClusterParameterGroup{
		Type:       "AWS::Redshift::ClusterParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRedshiftClusterParameterGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RedshiftClusterParameterGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftClusterParameterGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RedshiftClusterParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RedshiftClusterParameterGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.ParameterGroupFamily == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterGroupFamily'"))
	}
	return errs
}
