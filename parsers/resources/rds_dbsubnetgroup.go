package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSDBSubnetGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RDSDBSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSDBSubnetGroupProperties struct {
	DBSubnetGroupDescription interface{} `yaml:"DBSubnetGroupDescription"`
	DBSubnetGroupName interface{} `yaml:"DBSubnetGroupName,omitempty"`
	SubnetIds interface{} `yaml:"SubnetIds"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRDSDBSubnetGroup(properties RDSDBSubnetGroupProperties, deps ...interface{}) RDSDBSubnetGroup {
	return RDSDBSubnetGroup{
		Type:       "AWS::RDS::DBSubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSDBSubnetGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSDBSubnetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBSubnetGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSDBSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSDBSubnetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.DBSubnetGroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBSubnetGroupDescription'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
