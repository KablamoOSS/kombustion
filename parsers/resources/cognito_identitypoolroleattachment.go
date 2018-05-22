package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CognitoIdentityPoolRoleAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties CognitoIdentityPoolRoleAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CognitoIdentityPoolRoleAttachmentProperties struct {
	IdentityPoolId interface{} `yaml:"IdentityPoolId"`
	RoleMappings interface{} `yaml:"RoleMappings,omitempty"`
	Roles interface{} `yaml:"Roles,omitempty"`
}

func NewCognitoIdentityPoolRoleAttachment(properties CognitoIdentityPoolRoleAttachmentProperties, deps ...interface{}) CognitoIdentityPoolRoleAttachment {
	return CognitoIdentityPoolRoleAttachment{
		Type:       "AWS::Cognito::IdentityPoolRoleAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCognitoIdentityPoolRoleAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource CognitoIdentityPoolRoleAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoIdentityPoolRoleAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CognitoIdentityPoolRoleAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CognitoIdentityPoolRoleAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.IdentityPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IdentityPoolId'"))
	}
	return errs
}
