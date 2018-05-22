package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type KMSAlias struct {
	Type       string                      `yaml:"Type"`
	Properties KMSAliasProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type KMSAliasProperties struct {
	AliasName interface{} `yaml:"AliasName"`
	TargetKeyId interface{} `yaml:"TargetKeyId"`
}

func NewKMSAlias(properties KMSAliasProperties, deps ...interface{}) KMSAlias {
	return KMSAlias{
		Type:       "AWS::KMS::Alias",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKMSAlias(name string, data string) (cf types.ValueMap, err error) {
	var resource KMSAlias
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KMSAlias - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource KMSAlias) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KMSAliasProperties) Validate() []error {
	errs := []error{}
	if resource.AliasName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AliasName'"))
	}
	if resource.TargetKeyId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetKeyId'"))
	}
	return errs
}
