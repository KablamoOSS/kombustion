package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type Route53HealthCheck struct {
	Type       string                      `yaml:"Type"`
	Properties Route53HealthCheckProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type Route53HealthCheckProperties struct {
	HealthCheckTags interface{} `yaml:"HealthCheckTags,omitempty"`
	HealthCheckConfig *properties.HealthCheck_HealthCheckConfig `yaml:"HealthCheckConfig"`
}

func NewRoute53HealthCheck(properties Route53HealthCheckProperties, deps ...interface{}) Route53HealthCheck {
	return Route53HealthCheck{
		Type:       "AWS::Route53::HealthCheck",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRoute53HealthCheck(name string, data string) (cf types.ValueMap, err error) {
	var resource Route53HealthCheck
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Route53HealthCheck - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource Route53HealthCheck) Validate() []error {
	return resource.Properties.Validate()
}

func (resource Route53HealthCheckProperties) Validate() []error {
	errs := []error{}
	if resource.HealthCheckConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HealthCheckConfig'"))
	} else {
		errs = append(errs, resource.HealthCheckConfig.Validate()...)
	}
	return errs
}
