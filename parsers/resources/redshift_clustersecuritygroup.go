package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RedshiftClusterSecurityGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RedshiftClusterSecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RedshiftClusterSecurityGroupProperties struct {
	Description interface{} `yaml:"Description"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRedshiftClusterSecurityGroup(properties RedshiftClusterSecurityGroupProperties, deps ...interface{}) RedshiftClusterSecurityGroup {
	return RedshiftClusterSecurityGroup{
		Type:       "AWS::Redshift::ClusterSecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRedshiftClusterSecurityGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RedshiftClusterSecurityGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftClusterSecurityGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RedshiftClusterSecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RedshiftClusterSecurityGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	return errs
}
