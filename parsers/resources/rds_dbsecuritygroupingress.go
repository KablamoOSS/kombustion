package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSDBSecurityGroupIngress struct {
	Type       string                      `yaml:"Type"`
	Properties RDSDBSecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSDBSecurityGroupIngressProperties struct {
	CIDRIP interface{} `yaml:"CIDRIP,omitempty"`
	DBSecurityGroupName interface{} `yaml:"DBSecurityGroupName"`
	EC2SecurityGroupId interface{} `yaml:"EC2SecurityGroupId,omitempty"`
	EC2SecurityGroupName interface{} `yaml:"EC2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerId interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

func NewRDSDBSecurityGroupIngress(properties RDSDBSecurityGroupIngressProperties, deps ...interface{}) RDSDBSecurityGroupIngress {
	return RDSDBSecurityGroupIngress{
		Type:       "AWS::RDS::DBSecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSDBSecurityGroupIngress(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSDBSecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBSecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSDBSecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSDBSecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.DBSecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBSecurityGroupName'"))
	}
	return errs
}
