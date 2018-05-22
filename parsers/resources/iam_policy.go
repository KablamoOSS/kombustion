package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IAMPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties IAMPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMPolicyProperties struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	PolicyName interface{} `yaml:"PolicyName"`
	Groups interface{} `yaml:"Groups,omitempty"`
	Roles interface{} `yaml:"Roles,omitempty"`
	Users interface{} `yaml:"Users,omitempty"`
}

func NewIAMPolicy(properties IAMPolicyProperties, deps ...interface{}) IAMPolicy {
	return IAMPolicy{
		Type:       "AWS::IAM::Policy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	return errs
}
