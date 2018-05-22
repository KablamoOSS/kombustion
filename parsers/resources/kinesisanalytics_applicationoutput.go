package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type KinesisAnalyticsApplicationOutput struct {
	Type       string                      `yaml:"Type"`
	Properties KinesisAnalyticsApplicationOutputProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type KinesisAnalyticsApplicationOutputProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName"`
	Output *properties.ApplicationOutput_Output `yaml:"Output"`
}

func NewKinesisAnalyticsApplicationOutput(properties KinesisAnalyticsApplicationOutputProperties, deps ...interface{}) KinesisAnalyticsApplicationOutput {
	return KinesisAnalyticsApplicationOutput{
		Type:       "AWS::KinesisAnalytics::ApplicationOutput",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKinesisAnalyticsApplicationOutput(name string, data string) (cf types.ValueMap, err error) {
	var resource KinesisAnalyticsApplicationOutput
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisAnalyticsApplicationOutput - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource KinesisAnalyticsApplicationOutput) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KinesisAnalyticsApplicationOutputProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.Output == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Output'"))
	} else {
		errs = append(errs, resource.Output.Validate()...)
	}
	return errs
}
