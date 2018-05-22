package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSDBParameterGroup struct {
	Type       string                      `yaml:"Type"`
	Properties RDSDBParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSDBParameterGroupProperties struct {
	Description interface{} `yaml:"Description"`
	Family interface{} `yaml:"Family"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewRDSDBParameterGroup(properties RDSDBParameterGroupProperties, deps ...interface{}) RDSDBParameterGroup {
	return RDSDBParameterGroup{
		Type:       "AWS::RDS::DBParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSDBParameterGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSDBParameterGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBParameterGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSDBParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSDBParameterGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.Family == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Family'"))
	}
	return errs
}
