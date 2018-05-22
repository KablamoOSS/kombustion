package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CognitoUserPoolGroup struct {
	Type       string                      `yaml:"Type"`
	Properties CognitoUserPoolGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CognitoUserPoolGroupProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	GroupName interface{} `yaml:"GroupName,omitempty"`
	Precedence interface{} `yaml:"Precedence,omitempty"`
	RoleArn interface{} `yaml:"RoleArn,omitempty"`
	UserPoolId interface{} `yaml:"UserPoolId"`
}

func NewCognitoUserPoolGroup(properties CognitoUserPoolGroupProperties, deps ...interface{}) CognitoUserPoolGroup {
	return CognitoUserPoolGroup{
		Type:       "AWS::Cognito::UserPoolGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCognitoUserPoolGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource CognitoUserPoolGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CognitoUserPoolGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CognitoUserPoolGroupProperties) Validate() []error {
	errs := []error{}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	return errs
}
