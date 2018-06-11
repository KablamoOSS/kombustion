package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// AppSyncDataSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html
type AppSyncDataSource struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncDataSourceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// AppSyncDataSource Properties
type AppSyncDataSourceProperties struct {
	ApiId               interface{}                               `yaml:"ApiId"`
	Description         interface{}                               `yaml:"Description,omitempty"`
	Name                interface{}                               `yaml:"Name"`
	ServiceRoleArn      interface{}                               `yaml:"ServiceRoleArn,omitempty"`
	Type                interface{}                               `yaml:"Type"`
	LambdaConfig        *properties.DataSourceLambdaConfig        `yaml:"LambdaConfig,omitempty"`
	ElasticsearchConfig *properties.DataSourceElasticsearchConfig `yaml:"ElasticsearchConfig,omitempty"`
	DynamoDBConfig      *properties.DataSourceDynamoDBConfig      `yaml:"DynamoDBConfig,omitempty"`
}

// NewAppSyncDataSource constructor creates a new AppSyncDataSource
func NewAppSyncDataSource(properties AppSyncDataSourceProperties, deps ...interface{}) AppSyncDataSource {
	return AppSyncDataSource{
		Type:       "AWS::AppSync::DataSource",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppSyncDataSource parses AppSyncDataSource
func ParseAppSyncDataSource(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource AppSyncDataSource
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncDataSource - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource AppSyncDataSource) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AppSyncDataSourceProperties) Validate() []error {
	errs := []error{}
	if resource.ApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApiId'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
