package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueClassifier struct {
	Type       string                      `yaml:"Type"`
	Properties GlueClassifierProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueClassifierProperties struct {
	GrokClassifier *properties.Classifier_GrokClassifier `yaml:"GrokClassifier,omitempty"`
}

func NewGlueClassifier(properties GlueClassifierProperties, deps ...interface{}) GlueClassifier {
	return GlueClassifier{
		Type:       "AWS::Glue::Classifier",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueClassifier(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueClassifier
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueClassifier - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueClassifier) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueClassifierProperties) Validate() []error {
	errs := []error{}
	return errs
}
