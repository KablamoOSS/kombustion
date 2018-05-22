package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RedshiftClusterSecurityGroupIngress struct {
	Type       string                      `yaml:"Type"`
	Properties RedshiftClusterSecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RedshiftClusterSecurityGroupIngressProperties struct {
	CIDRIP interface{} `yaml:"CIDRIP,omitempty"`
	ClusterSecurityGroupName interface{} `yaml:"ClusterSecurityGroupName"`
	EC2SecurityGroupName interface{} `yaml:"EC2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerId interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

func NewRedshiftClusterSecurityGroupIngress(properties RedshiftClusterSecurityGroupIngressProperties, deps ...interface{}) RedshiftClusterSecurityGroupIngress {
	return RedshiftClusterSecurityGroupIngress{
		Type:       "AWS::Redshift::ClusterSecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRedshiftClusterSecurityGroupIngress(name string, data string) (cf types.ValueMap, err error) {
	var resource RedshiftClusterSecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftClusterSecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RedshiftClusterSecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RedshiftClusterSecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.ClusterSecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ClusterSecurityGroupName'"))
	}
	return errs
}
