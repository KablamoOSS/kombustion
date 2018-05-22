package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type SSMDocument struct {
	Type       string                      `yaml:"Type"`
	Properties SSMDocumentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SSMDocumentProperties struct {
	Content interface{} `yaml:"Content"`
	DocumentType interface{} `yaml:"DocumentType,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewSSMDocument(properties SSMDocumentProperties, deps ...interface{}) SSMDocument {
	return SSMDocument{
		Type:       "AWS::SSM::Document",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSSMDocument(name string, data string) (cf types.ValueMap, err error) {
	var resource SSMDocument
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMDocument - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SSMDocument) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SSMDocumentProperties) Validate() []error {
	errs := []error{}
	if resource.Content == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Content'"))
	}
	return errs
}
