package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type KMSKey struct {
	Type       string                      `yaml:"Type"`
	Properties KMSKeyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type KMSKeyProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	EnableKeyRotation interface{} `yaml:"EnableKeyRotation,omitempty"`
	Enabled interface{} `yaml:"Enabled,omitempty"`
	KeyPolicy interface{} `yaml:"KeyPolicy"`
	KeyUsage interface{} `yaml:"KeyUsage,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewKMSKey(properties KMSKeyProperties, deps ...interface{}) KMSKey {
	return KMSKey{
		Type:       "AWS::KMS::Key",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKMSKey(name string, data string) (cf types.ValueMap, err error) {
	var resource KMSKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KMSKey - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource KMSKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KMSKeyProperties) Validate() []error {
	errs := []error{}
	if resource.KeyPolicy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyPolicy'"))
	}
	return errs
}
