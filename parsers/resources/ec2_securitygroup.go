package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2SecurityGroup struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SecurityGroupProperties struct {
	GroupDescription interface{} `yaml:"GroupDescription"`
	GroupName interface{} `yaml:"GroupName,omitempty"`
	VpcId interface{} `yaml:"VpcId,omitempty"`
	SecurityGroupEgress interface{} `yaml:"SecurityGroupEgress,omitempty"`
	SecurityGroupIngress interface{} `yaml:"SecurityGroupIngress,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2SecurityGroup(properties EC2SecurityGroupProperties, deps ...interface{}) EC2SecurityGroup {
	return EC2SecurityGroup{
		Type:       "AWS::EC2::SecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SecurityGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2SecurityGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SecurityGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2SecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SecurityGroupProperties) Validate() []error {
	errs := []error{}
	if resource.GroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupDescription'"))
	}
	return errs
}
