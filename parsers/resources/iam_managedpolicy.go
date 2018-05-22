package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IAMManagedPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties IAMManagedPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMManagedPolicyProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	ManagedPolicyName interface{} `yaml:"ManagedPolicyName,omitempty"`
	Path interface{} `yaml:"Path,omitempty"`
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	Groups interface{} `yaml:"Groups,omitempty"`
	Roles interface{} `yaml:"Roles,omitempty"`
	Users interface{} `yaml:"Users,omitempty"`
}

func NewIAMManagedPolicy(properties IAMManagedPolicyProperties, deps ...interface{}) IAMManagedPolicy {
	return IAMManagedPolicy{
		Type:       "AWS::IAM::ManagedPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMManagedPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMManagedPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMManagedPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMManagedPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMManagedPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	return errs
}
