package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSDBSecurityGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RDSDBSecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSDBSecurityGroupProperties struct {
	EC2VpcId interface{} `yaml:"EC2VpcId,omitempty"`
	GroupDescription interface{} `yaml:"GroupDescription"`
	DBSecurityGroupIngress interface{} `yaml:"DBSecurityGroupIngress"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRDSDBSecurityGroup(properties RDSDBSecurityGroupProperties, deps ...interface{}) RDSDBSecurityGroup {
	return RDSDBSecurityGroup{
		Type:       "AWS::RDS::DBSecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSDBSecurityGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSDBSecurityGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBSecurityGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSDBSecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSDBSecurityGroupProperties) Validate() []error {
	errs := []error{}
	if resource.GroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupDescription'"))
	}
	if resource.DBSecurityGroupIngress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBSecurityGroupIngress'"))
	}
	return errs
}
