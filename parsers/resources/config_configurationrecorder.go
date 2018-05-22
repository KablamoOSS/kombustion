package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ConfigConfigurationRecorder struct {
	Type       string                      `yaml:"Type"`
	Properties ConfigConfigurationRecorderProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ConfigConfigurationRecorderProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	RoleARN interface{} `yaml:"RoleARN"`
	RecordingGroup *properties.ConfigurationRecorder_RecordingGroup `yaml:"RecordingGroup,omitempty"`
}

func NewConfigConfigurationRecorder(properties ConfigConfigurationRecorderProperties, deps ...interface{}) ConfigConfigurationRecorder {
	return ConfigConfigurationRecorder{
		Type:       "AWS::Config::ConfigurationRecorder",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseConfigConfigurationRecorder(name string, data string) (cf types.ValueMap, err error) {
	var resource ConfigConfigurationRecorder
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ConfigConfigurationRecorder - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ConfigConfigurationRecorder) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ConfigConfigurationRecorderProperties) Validate() []error {
	errs := []error{}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errs
}
