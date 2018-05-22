package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IAMRole struct {
	Type       string                      `yaml:"Type"`
	Properties IAMRoleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMRoleProperties struct {
	AssumeRolePolicyDocument interface{} `yaml:"AssumeRolePolicyDocument"`
	Path interface{} `yaml:"Path,omitempty"`
	RoleName interface{} `yaml:"RoleName,omitempty"`
	ManagedPolicyArns interface{} `yaml:"ManagedPolicyArns,omitempty"`
	Policies interface{} `yaml:"Policies,omitempty"`
}

func NewIAMRole(properties IAMRoleProperties, deps ...interface{}) IAMRole {
	return IAMRole{
		Type:       "AWS::IAM::Role",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMRole(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMRole
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMRole - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMRole) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMRoleProperties) Validate() []error {
	errs := []error{}
	if resource.AssumeRolePolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AssumeRolePolicyDocument'"))
	}
	return errs
}
