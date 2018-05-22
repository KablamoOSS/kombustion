package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type Route53HostedZone struct {
	Type       string                      `yaml:"Type"`
	Properties Route53HostedZoneProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type Route53HostedZoneProperties struct {
	Name interface{} `yaml:"Name"`
	QueryLoggingConfig *properties.HostedZone_QueryLoggingConfig `yaml:"QueryLoggingConfig,omitempty"`
	HostedZoneTags interface{} `yaml:"HostedZoneTags,omitempty"`
	VPCs interface{} `yaml:"VPCs,omitempty"`
	HostedZoneConfig *properties.HostedZone_HostedZoneConfig `yaml:"HostedZoneConfig,omitempty"`
}

func NewRoute53HostedZone(properties Route53HostedZoneProperties, deps ...interface{}) Route53HostedZone {
	return Route53HostedZone{
		Type:       "AWS::Route53::HostedZone",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRoute53HostedZone(name string, data string) (cf types.ValueMap, err error) {
	var resource Route53HostedZone
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Route53HostedZone - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource Route53HostedZone) Validate() []error {
	return resource.Properties.Validate()
}

func (resource Route53HostedZoneProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
