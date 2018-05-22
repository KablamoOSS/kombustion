package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ElasticBeanstalkConfigurationTemplate struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticBeanstalkConfigurationTemplateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticBeanstalkConfigurationTemplateProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName"`
	Description interface{} `yaml:"Description,omitempty"`
	EnvironmentId interface{} `yaml:"EnvironmentId,omitempty"`
	PlatformArn interface{} `yaml:"PlatformArn,omitempty"`
	SolutionStackName interface{} `yaml:"SolutionStackName,omitempty"`
	SourceConfiguration *properties.ConfigurationTemplate_SourceConfiguration `yaml:"SourceConfiguration,omitempty"`
	OptionSettings interface{} `yaml:"OptionSettings,omitempty"`
}

func NewElasticBeanstalkConfigurationTemplate(properties ElasticBeanstalkConfigurationTemplateProperties, deps ...interface{}) ElasticBeanstalkConfigurationTemplate {
	return ElasticBeanstalkConfigurationTemplate{
		Type:       "AWS::ElasticBeanstalk::ConfigurationTemplate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticBeanstalkConfigurationTemplate(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticBeanstalkConfigurationTemplate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkConfigurationTemplate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticBeanstalkConfigurationTemplate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticBeanstalkConfigurationTemplateProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	return errs
}
