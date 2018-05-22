package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IAMInstanceProfile struct {
	Type       string                      `yaml:"Type"`
	Properties IAMInstanceProfileProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMInstanceProfileProperties struct {
	InstanceProfileName interface{} `yaml:"InstanceProfileName,omitempty"`
	Path interface{} `yaml:"Path,omitempty"`
	Roles interface{} `yaml:"Roles"`
}

func NewIAMInstanceProfile(properties IAMInstanceProfileProperties, deps ...interface{}) IAMInstanceProfile {
	return IAMInstanceProfile{
		Type:       "AWS::IAM::InstanceProfile",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMInstanceProfile(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMInstanceProfile
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMInstanceProfile - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMInstanceProfile) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMInstanceProfileProperties) Validate() []error {
	errs := []error{}
	if resource.Roles == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Roles'"))
	}
	return errs
}
