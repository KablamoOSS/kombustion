package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EMRSecurityConfiguration struct {
	Type       string                      `yaml:"Type"`
	Properties EMRSecurityConfigurationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EMRSecurityConfigurationProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	SecurityConfiguration interface{} `yaml:"SecurityConfiguration"`
}

func NewEMRSecurityConfiguration(properties EMRSecurityConfigurationProperties, deps ...interface{}) EMRSecurityConfiguration {
	return EMRSecurityConfiguration{
		Type:       "AWS::EMR::SecurityConfiguration",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEMRSecurityConfiguration(name string, data string) (cf types.ValueMap, err error) {
	var resource EMRSecurityConfiguration
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRSecurityConfiguration - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EMRSecurityConfiguration) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EMRSecurityConfigurationProperties) Validate() []error {
	errs := []error{}
	if resource.SecurityConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityConfiguration'"))
	}
	return errs
}
