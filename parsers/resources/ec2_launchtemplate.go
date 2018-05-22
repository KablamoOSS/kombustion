package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EC2LaunchTemplate struct {
	Type       string                      `yaml:"Type"`
	Properties EC2LaunchTemplateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2LaunchTemplateProperties struct {
	LaunchTemplateName interface{} `yaml:"LaunchTemplateName,omitempty"`
	LaunchTemplateData *properties.LaunchTemplate_LaunchTemplateData `yaml:"LaunchTemplateData,omitempty"`
}

func NewEC2LaunchTemplate(properties EC2LaunchTemplateProperties, deps ...interface{}) EC2LaunchTemplate {
	return EC2LaunchTemplate{
		Type:       "AWS::EC2::LaunchTemplate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2LaunchTemplate(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2LaunchTemplate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2LaunchTemplate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2LaunchTemplate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2LaunchTemplateProperties) Validate() []error {
	errs := []error{}
	return errs
}
