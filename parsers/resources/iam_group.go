package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type IAMGroup struct {
	Type       string                      `yaml:"Type"`
	Properties IAMGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMGroupProperties struct {
	GroupName interface{} `yaml:"GroupName,omitempty"`
	Path interface{} `yaml:"Path,omitempty"`
	ManagedPolicyArns interface{} `yaml:"ManagedPolicyArns,omitempty"`
	Policies interface{} `yaml:"Policies,omitempty"`
}

func NewIAMGroup(properties IAMGroupProperties, deps ...interface{}) IAMGroup {
	return IAMGroup{
		Type:       "AWS::IAM::Group",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
