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

// KinesisAnalyticsApplicationOutput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html
type KinesisAnalyticsApplicationOutput struct {
	Type       string                                      `yaml:"Type"`
	Properties KinesisAnalyticsApplicationOutputProperties `yaml:"Properties"`
	Condition  interface{}                                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                 `yaml:"DependsOn,omitempty"`
}

// KinesisAnalyticsApplicationOutput Properties
type KinesisAnalyticsApplicationOutputProperties struct {
	ApplicationName interface{}                         `yaml:"ApplicationName"`
	Output          *properties.ApplicationOutputOutput `yaml:"Output"`
}

// NewKinesisAnalyticsApplicationOutput constructor creates a new KinesisAnalyticsApplicationOutput
func NewKinesisAnalyticsApplicationOutput(properties KinesisAnalyticsApplicationOutputProperties, deps ...interface{}) KinesisAnalyticsApplicationOutput {
	return KinesisAnalyticsApplicationOutput{
		Type:       "AWS::KinesisAnalytics::ApplicationOutput",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseKinesisAnalyticsApplicationOutput parses KinesisAnalyticsApplicationOutput
func ParseKinesisAnalyticsApplicationOutput(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
