package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CognitoUserPoolUser struct {
	Type       string                      `yaml:"Type"`
	Properties CognitoUserPoolUserProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CognitoUserPoolUserProperties struct {
	ForceAliasCreation interface{} `yaml:"ForceAliasCreation,omitempty"`
	MessageAction interface{} `yaml:"MessageAction,omitempty"`
	UserPoolId interface{} `yaml:"UserPoolId"`
	Username interface{} `yaml:"Username,omitempty"`
	DesiredDeliveryMediums interface{} `yaml:"DesiredDeliveryMediums,omitempty"`
	UserAttributes interface{} `yaml:"UserAttributes,omitempty"`
	ValidationData interface{} `yaml:"ValidationData,omitempty"`
}

func NewCognitoUserPoolUser(properties CognitoUserPoolUserProperties, deps ...interface{}) CognitoUserPoolUser {
	return CognitoUserPoolUser{
		Type:       "AWS::Cognito::UserPoolUser",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCognitoUserPoolUser(name string, data string) (cf types.ValueMap, err error) {
	var resource CognitoUserPoolUser
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolUser - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CognitoUserPoolUser) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CognitoUserPoolUserProperties) Validate() []error {
	errs := []error{}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	return errs
}
