package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CognitoIdentityPool struct {
	Type       string                      `yaml:"Type"`
	Properties CognitoIdentityPoolProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CognitoIdentityPoolProperties struct {
	AllowUnauthenticatedIdentities interface{} `yaml:"AllowUnauthenticatedIdentities"`
	CognitoEvents interface{} `yaml:"CognitoEvents,omitempty"`
	DeveloperProviderName interface{} `yaml:"DeveloperProviderName,omitempty"`
	IdentityPoolName interface{} `yaml:"IdentityPoolName,omitempty"`
	SupportedLoginProviders interface{} `yaml:"SupportedLoginProviders,omitempty"`
	PushSync *properties.IdentityPool_PushSync `yaml:"PushSync,omitempty"`
	CognitoIdentityProviders interface{} `yaml:"CognitoIdentityProviders,omitempty"`
	OpenIdConnectProviderARNs interface{} `yaml:"OpenIdConnectProviderARNs,omitempty"`
	SamlProviderARNs interface{} `yaml:"SamlProviderARNs,omitempty"`
	CognitoStreams *properties.IdentityPool_CognitoStreams `yaml:"CognitoStreams,omitempty"`
}

func NewCognitoIdentityPool(properties CognitoIdentityPoolProperties, deps ...interface{}) CognitoIdentityPool {
	return CognitoIdentityPool{
		Type:       "AWS::Cognito::IdentityPool",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCognitoIdentityPool(name string, data string) (cf types.ValueMap, err error) {
	var resource CognitoIdentityPool
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoIdentityPool - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CognitoIdentityPool) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CognitoIdentityPoolProperties) Validate() []error {
	errs := []error{}
	if resource.AllowUnauthenticatedIdentities == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AllowUnauthenticatedIdentities'"))
	}
	return errs
}
