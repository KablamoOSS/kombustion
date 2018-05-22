package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type SESConfigurationSet struct {
	Type       string                      `yaml:"Type"`
	Properties SESConfigurationSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SESConfigurationSetProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
}

func NewSESConfigurationSet(properties SESConfigurationSetProperties, deps ...interface{}) SESConfigurationSet {
	return SESConfigurationSet{
		Type:       "AWS::SES::ConfigurationSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESConfigurationSet(name string, data string) (cf types.ValueMap, err error) {
	var resource SESConfigurationSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESConfigurationSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SESConfigurationSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESConfigurationSetProperties) Validate() []error {
	errs := []error{}
	return errs
}
