package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type OpsWorksUserProfile struct {
	Type       string                      `yaml:"Type"`
	Properties OpsWorksUserProfileProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type OpsWorksUserProfileProperties struct {
	AllowSelfManagement interface{} `yaml:"AllowSelfManagement,omitempty"`
	IamUserArn interface{} `yaml:"IamUserArn"`
	SshPublicKey interface{} `yaml:"SshPublicKey,omitempty"`
	SshUsername interface{} `yaml:"SshUsername,omitempty"`
}

func NewOpsWorksUserProfile(properties OpsWorksUserProfileProperties, deps ...interface{}) OpsWorksUserProfile {
	return OpsWorksUserProfile{
		Type:       "AWS::OpsWorks::UserProfile",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksUserProfile(name string, data string) (cf types.ValueMap, err error) {
	var resource OpsWorksUserProfile
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksUserProfile - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource OpsWorksUserProfile) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksUserProfileProperties) Validate() []error {
	errs := []error{}
	if resource.IamUserArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IamUserArn'"))
	}
	return errs
}
