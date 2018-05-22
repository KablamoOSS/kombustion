package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CloudFormationCustomResource struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFormationCustomResourceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFormationCustomResourceProperties struct {
	ServiceToken interface{} `yaml:"ServiceToken"`
}

func NewCloudFormationCustomResource(properties CloudFormationCustomResourceProperties, deps ...interface{}) CloudFormationCustomResource {
	return CloudFormationCustomResource{
		Type:       "AWS::CloudFormation::CustomResource",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFormationCustomResource(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFormationCustomResource
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFormationCustomResource - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFormationCustomResource) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFormationCustomResourceProperties) Validate() []error {
	errs := []error{}
	if resource.ServiceToken == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceToken'"))
	}
	return errs
}
