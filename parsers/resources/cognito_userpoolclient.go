package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CognitoUserPoolClient struct {
	Type       string                      `yaml:"Type"`
	Properties CognitoUserPoolClientProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CognitoUserPoolClientProperties struct {
	ClientName interface{} `yaml:"ClientName,omitempty"`
	GenerateSecret interface{} `yaml:"GenerateSecret,omitempty"`
	RefreshTokenValidity interface{} `yaml:"RefreshTokenValidity,omitempty"`
	UserPoolId interface{} `yaml:"UserPoolId"`
	ExplicitAuthFlows interface{} `yaml:"ExplicitAuthFlows,omitempty"`
	ReadAttributes interface{} `yaml:"ReadAttributes,omitempty"`
	WriteAttributes interface{} `yaml:"WriteAttributes,omitempty"`
}

func NewCognitoUserPoolClient(properties CognitoUserPoolClientProperties, deps ...interface{}) CognitoUserPoolClient {
	return CognitoUserPoolClient{
		Type:       "AWS::Cognito::UserPoolClient",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCognitoUserPoolClient(name string, data string) (cf types.ValueMap, err error) {
	var resource CognitoUserPoolClient
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolClient - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CognitoUserPoolClient) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CognitoUserPoolClientProperties) Validate() []error {
	errs := []error{}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	return errs
}
