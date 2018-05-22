package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CognitoUserPoolUserToGroupAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties CognitoUserPoolUserToGroupAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CognitoUserPoolUserToGroupAttachmentProperties struct {
	GroupName interface{} `yaml:"GroupName"`
	UserPoolId interface{} `yaml:"UserPoolId"`
	Username interface{} `yaml:"Username"`
}

func NewCognitoUserPoolUserToGroupAttachment(properties CognitoUserPoolUserToGroupAttachmentProperties, deps ...interface{}) CognitoUserPoolUserToGroupAttachment {
	return CognitoUserPoolUserToGroupAttachment{
		Type:       "AWS::Cognito::UserPoolUserToGroupAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCognitoUserPoolUserToGroupAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource CognitoUserPoolUserToGroupAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolUserToGroupAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CognitoUserPoolUserToGroupAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CognitoUserPoolUserToGroupAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.GroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupName'"))
	}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	if resource.Username == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Username'"))
	}
	return errs
}
