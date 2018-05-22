package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IAMAccessKey struct {
	Type       string                      `yaml:"Type"`
	Properties IAMAccessKeyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IAMAccessKeyProperties struct {
	Serial interface{} `yaml:"Serial,omitempty"`
	Status interface{} `yaml:"Status,omitempty"`
	UserName interface{} `yaml:"UserName"`
}

func NewIAMAccessKey(properties IAMAccessKeyProperties, deps ...interface{}) IAMAccessKey {
	return IAMAccessKey{
		Type:       "AWS::IAM::AccessKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMAccessKey(name string, data string) (cf types.ValueMap, err error) {
	var resource IAMAccessKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMAccessKey - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IAMAccessKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMAccessKeyProperties) Validate() []error {
	errs := []error{}
	if resource.UserName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserName'"))
	}
	return errs
}
