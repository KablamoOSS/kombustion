package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type IAMUser struct {
	Type       string                      `yaml:"Type"`
	Properties IAMUserProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMUserProperties struct {
	Path interface{} `yaml:"Path,omitempty"`
	UserName interface{} `yaml:"UserName,omitempty"`
	LoginProfile *properties.User_LoginProfile `yaml:"LoginProfile,omitempty"`
	Groups interface{} `yaml:"Groups,omitempty"`
	ManagedPolicyArns interface{} `yaml:"ManagedPolicyArns,omitempty"`
	Policies interface{} `yaml:"Policies,omitempty"`
}

func NewIAMUser(properties IAMUserProperties, deps ...interface{}) IAMUser {
	return IAMUser{
		Type:       "AWS::IAM::User",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMUser(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMUser
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMUser - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMUser) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMUserProperties) Validate() []error {
	errs := []error{}
	return errs
}
