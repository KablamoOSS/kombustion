package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ElasticBeanstalkApplication struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticBeanstalkApplicationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticBeanstalkApplicationProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	ResourceLifecycleConfig *properties.Application_ApplicationResourceLifecycleConfig `yaml:"ResourceLifecycleConfig,omitempty"`
}

func NewElasticBeanstalkApplication(properties ElasticBeanstalkApplicationProperties, deps ...interface{}) ElasticBeanstalkApplication {
	return ElasticBeanstalkApplication{
		Type:       "AWS::ElasticBeanstalk::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticBeanstalkApplication(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticBeanstalkApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkApplication - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticBeanstalkApplication) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticBeanstalkApplicationProperties) Validate() []error {
	errs := []error{}
	return errs
}
