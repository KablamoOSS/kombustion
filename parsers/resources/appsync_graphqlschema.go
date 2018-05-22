package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AppSyncGraphQLSchema struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncGraphQLSchemaProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AppSyncGraphQLSchemaProperties struct {
	ApiId interface{} `yaml:"ApiId"`
	Definition interface{} `yaml:"Definition,omitempty"`
	DefinitionS3Location interface{} `yaml:"DefinitionS3Location,omitempty"`
}

func NewAppSyncGraphQLSchema(properties AppSyncGraphQLSchemaProperties, deps ...interface{}) AppSyncGraphQLSchema {
	return AppSyncGraphQLSchema{
		Type:       "AWS::AppSync::GraphQLSchema",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAppSyncGraphQLSchema(name string, data string) (cf types.ValueMap, err error) {
	var resource AppSyncGraphQLSchema
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncGraphQLSchema - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AppSyncGraphQLSchema) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AppSyncGraphQLSchemaProperties) Validate() []error {
	errs := []error{}
	if resource.ApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApiId'"))
	}
	return errs
}
