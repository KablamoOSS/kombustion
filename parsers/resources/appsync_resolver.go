package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AppSyncResolver struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncResolverProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AppSyncResolverProperties struct {
	ApiId interface{} `yaml:"ApiId"`
	DataSourceName interface{} `yaml:"DataSourceName"`
	FieldName interface{} `yaml:"FieldName"`
	RequestMappingTemplate interface{} `yaml:"RequestMappingTemplate,omitempty"`
	RequestMappingTemplateS3Location interface{} `yaml:"RequestMappingTemplateS3Location,omitempty"`
	ResponseMappingTemplate interface{} `yaml:"ResponseMappingTemplate,omitempty"`
	ResponseMappingTemplateS3Location interface{} `yaml:"ResponseMappingTemplateS3Location,omitempty"`
	TypeName interface{} `yaml:"TypeName"`
}

func NewAppSyncResolver(properties AppSyncResolverProperties, deps ...interface{}) AppSyncResolver {
	return AppSyncResolver{
		Type:       "AWS::AppSync::Resolver",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAppSyncResolver(name string, data string) (cf types.ValueMap, err error) {
	var resource AppSyncResolver
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncResolver - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AppSyncResolver) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AppSyncResolverProperties) Validate() []error {
	errs := []error{}
	if resource.ApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApiId'"))
	}
	if resource.DataSourceName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DataSourceName'"))
	}
	if resource.FieldName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FieldName'"))
	}
	if resource.TypeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TypeName'"))
	}
	return errs
}
