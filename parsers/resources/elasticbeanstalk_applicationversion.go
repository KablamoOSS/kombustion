package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ElasticBeanstalkApplicationVersion struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticBeanstalkApplicationVersionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticBeanstalkApplicationVersionProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName"`
	Description interface{} `yaml:"Description,omitempty"`
	SourceBundle *properties.ApplicationVersion_SourceBundle `yaml:"SourceBundle"`
}

func NewElasticBeanstalkApplicationVersion(properties ElasticBeanstalkApplicationVersionProperties, deps ...interface{}) ElasticBeanstalkApplicationVersion {
	return ElasticBeanstalkApplicationVersion{
		Type:       "AWS::ElasticBeanstalk::ApplicationVersion",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticBeanstalkApplicationVersion(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticBeanstalkApplicationVersion
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkApplicationVersion - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticBeanstalkApplicationVersion) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticBeanstalkApplicationVersionProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.SourceBundle == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SourceBundle'"))
	} else {
		errs = append(errs, resource.SourceBundle.Validate()...)
	}
	return errs
}
