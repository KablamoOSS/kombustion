package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CognitoUserPoolUserToGroupAttachment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html
type CognitoUserPoolUserToGroupAttachment struct {
	Type       string                                         `yaml:"Type"`
	Properties CognitoUserPoolUserToGroupAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                    `yaml:"DependsOn,omitempty"`
}

// CognitoUserPoolUserToGroupAttachment Properties
type CognitoUserPoolUserToGroupAttachmentProperties struct {
	GroupName  interface{} `yaml:"GroupName"`
	UserPoolId interface{} `yaml:"UserPoolId"`
	Username   interface{} `yaml:"Username"`
}

// NewCognitoUserPoolUserToGroupAttachment constructor creates a new CognitoUserPoolUserToGroupAttachment
func NewCognitoUserPoolUserToGroupAttachment(properties CognitoUserPoolUserToGroupAttachmentProperties, deps ...interface{}) CognitoUserPoolUserToGroupAttachment {
	return CognitoUserPoolUserToGroupAttachment{
		Type:       "AWS::Cognito::UserPoolUserToGroupAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoUserPoolUserToGroupAttachment parses CognitoUserPoolUserToGroupAttachment
func ParseCognitoUserPoolUserToGroupAttachment(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
