package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type AppSyncDataSource struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncDataSourceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AppSyncDataSourceProperties struct {
	ApiId interface{} `yaml:"ApiId"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	ServiceRoleArn interface{} `yaml:"ServiceRoleArn,omitempty"`
	Type interface{} `yaml:"Type"`
	LambdaConfig *properties.DataSource_LambdaConfig `yaml:"LambdaConfig,omitempty"`
	ElasticsearchConfig *properties.DataSource_ElasticsearchConfig `yaml:"ElasticsearchConfig,omitempty"`
	DynamoDBConfig *properties.DataSource_DynamoDBConfig `yaml:"DynamoDBConfig,omitempty"`
}

func NewAppSyncDataSource(properties AppSyncDataSourceProperties, deps ...interface{}) AppSyncDataSource {
	return AppSyncDataSource{
		Type:       "AWS::AppSync::DataSource",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAppSyncDataSource(name string, data string) (cf types.ValueMap, err error) {
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
	cf = types.ValueMap{name: resource}
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
