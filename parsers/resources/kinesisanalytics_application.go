package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type KinesisAnalyticsApplication struct {
	Type       string                      `yaml:"Type"`
	Properties KinesisAnalyticsApplicationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type KinesisAnalyticsApplicationProperties struct {
	ApplicationCode interface{} `yaml:"ApplicationCode,omitempty"`
	ApplicationDescription interface{} `yaml:"ApplicationDescription,omitempty"`
	ApplicationName interface{} `yaml:"ApplicationName,omitempty"`
	Inputs interface{} `yaml:"Inputs"`
}

func NewKinesisAnalyticsApplication(properties KinesisAnalyticsApplicationProperties, deps ...interface{}) KinesisAnalyticsApplication {
	return KinesisAnalyticsApplication{
		Type:       "AWS::KinesisAnalytics::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKinesisAnalyticsApplication(name string, data string) (cf types.ValueMap, err error) {
	var resource KinesisAnalyticsApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisAnalyticsApplication - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource KinesisAnalyticsApplication) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KinesisAnalyticsApplicationProperties) Validate() []error {
	errs := []error{}
	if resource.Inputs == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Inputs'"))
	}
	return errs
}
