package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IoTPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties IoTPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IoTPolicyProperties struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	PolicyName interface{} `yaml:"PolicyName,omitempty"`
}

func NewIoTPolicy(properties IoTPolicyProperties, deps ...interface{}) IoTPolicy {
	return IoTPolicy{
		Type:       "AWS::IoT::Policy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource IoTPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IoTPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	return errs
}
