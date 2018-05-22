package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type SSMAssociation struct {
	Type       string                      `yaml:"Type"`
	Properties SSMAssociationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SSMAssociationProperties struct {
	AssociationName interface{} `yaml:"AssociationName,omitempty"`
	DocumentVersion interface{} `yaml:"DocumentVersion,omitempty"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
	Name interface{} `yaml:"Name"`
	ScheduleExpression interface{} `yaml:"ScheduleExpression,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	Targets interface{} `yaml:"Targets,omitempty"`
}

func NewSSMAssociation(properties SSMAssociationProperties, deps ...interface{}) SSMAssociation {
	return SSMAssociation{
		Type:       "AWS::SSM::Association",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSSMAssociation(name string, data string) (cf types.ValueMap, err error) {
	var resource SSMAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMAssociation - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SSMAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SSMAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
