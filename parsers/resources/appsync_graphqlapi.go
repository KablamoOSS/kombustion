package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type AppSyncGraphQLApi struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncGraphQLApiProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AppSyncGraphQLApiProperties struct {
	AuthenticationType interface{} `yaml:"AuthenticationType"`
	Name interface{} `yaml:"Name"`
	UserPoolConfig *properties.GraphQLApi_UserPoolConfig `yaml:"UserPoolConfig,omitempty"`
	OpenIDConnectConfig *properties.GraphQLApi_OpenIDConnectConfig `yaml:"OpenIDConnectConfig,omitempty"`
	LogConfig *properties.GraphQLApi_LogConfig `yaml:"LogConfig,omitempty"`
}

func NewAppSyncGraphQLApi(properties AppSyncGraphQLApiProperties, deps ...interface{}) AppSyncGraphQLApi {
	return AppSyncGraphQLApi{
		Type:       "AWS::AppSync::GraphQLApi",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAppSyncGraphQLApi(name string, data string) (cf types.ValueMap, err error) {
	var resource AppSyncGraphQLApi
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncGraphQLApi - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AppSyncGraphQLApi) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AppSyncGraphQLApiProperties) Validate() []error {
	errs := []error{}
	if resource.AuthenticationType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AuthenticationType'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
