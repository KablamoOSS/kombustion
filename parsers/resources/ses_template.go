package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type SESTemplate struct {
	Type       string                      `yaml:"Type"`
	Properties SESTemplateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SESTemplateProperties struct {
	Template *properties.Template_Template `yaml:"Template,omitempty"`
}

func NewSESTemplate(properties SESTemplateProperties, deps ...interface{}) SESTemplate {
	return SESTemplate{
		Type:       "AWS::SES::Template",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESTemplate(name string, data string) (cf types.ValueMap, err error) {
	var resource SESTemplate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESTemplate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SESTemplate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESTemplateProperties) Validate() []error {
	errs := []error{}
	return errs
}
