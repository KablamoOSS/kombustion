package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type KinesisAnalyticsApplicationReferenceDataSource struct {
	Type       string                      `yaml:"Type"`
	Properties KinesisAnalyticsApplicationReferenceDataSourceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type KinesisAnalyticsApplicationReferenceDataSourceProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName"`
	ReferenceDataSource *properties.ApplicationReferenceDataSource_ReferenceDataSource `yaml:"ReferenceDataSource"`
}

func NewKinesisAnalyticsApplicationReferenceDataSource(properties KinesisAnalyticsApplicationReferenceDataSourceProperties, deps ...interface{}) KinesisAnalyticsApplicationReferenceDataSource {
	return KinesisAnalyticsApplicationReferenceDataSource{
		Type:       "AWS::KinesisAnalytics::ApplicationReferenceDataSource",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKinesisAnalyticsApplicationReferenceDataSource(name string, data string) (cf types.ValueMap, err error) {
	var resource KinesisAnalyticsApplicationReferenceDataSource
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisAnalyticsApplicationReferenceDataSource - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource KinesisAnalyticsApplicationReferenceDataSource) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KinesisAnalyticsApplicationReferenceDataSourceProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.ReferenceDataSource == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReferenceDataSource'"))
	} else {
		errs = append(errs, resource.ReferenceDataSource.Validate()...)
	}
	return errs
}
