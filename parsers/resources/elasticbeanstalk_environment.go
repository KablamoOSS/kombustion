package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ElasticBeanstalkEnvironment struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticBeanstalkEnvironmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticBeanstalkEnvironmentProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName"`
	CNAMEPrefix interface{} `yaml:"CNAMEPrefix,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	EnvironmentName interface{} `yaml:"EnvironmentName,omitempty"`
	PlatformArn interface{} `yaml:"PlatformArn,omitempty"`
	SolutionStackName interface{} `yaml:"SolutionStackName,omitempty"`
	TemplateName interface{} `yaml:"TemplateName,omitempty"`
	VersionLabel interface{} `yaml:"VersionLabel,omitempty"`
	Tier *properties.Environment_Tier `yaml:"Tier,omitempty"`
	OptionSettings interface{} `yaml:"OptionSettings,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewElasticBeanstalkEnvironment(properties ElasticBeanstalkEnvironmentProperties, deps ...interface{}) ElasticBeanstalkEnvironment {
	return ElasticBeanstalkEnvironment{
		Type:       "AWS::ElasticBeanstalk::Environment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticBeanstalkEnvironment(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticBeanstalkEnvironment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkEnvironment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticBeanstalkEnvironment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticBeanstalkEnvironmentProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	return errs
}
