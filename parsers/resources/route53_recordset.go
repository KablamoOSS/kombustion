package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type Route53RecordSet struct {
	Type       string                      `yaml:"Type"`
	Properties Route53RecordSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type Route53RecordSetProperties struct {
	Comment interface{} `yaml:"Comment,omitempty"`
	Failover interface{} `yaml:"Failover,omitempty"`
	HealthCheckId interface{} `yaml:"HealthCheckId,omitempty"`
	HostedZoneId interface{} `yaml:"HostedZoneId,omitempty"`
	HostedZoneName interface{} `yaml:"HostedZoneName,omitempty"`
	Name interface{} `yaml:"Name"`
	Region interface{} `yaml:"Region,omitempty"`
	SetIdentifier interface{} `yaml:"SetIdentifier,omitempty"`
	TTL interface{} `yaml:"TTL,omitempty"`
	Type interface{} `yaml:"Type"`
	Weight interface{} `yaml:"Weight,omitempty"`
	ResourceRecords interface{} `yaml:"ResourceRecords,omitempty"`
	GeoLocation *properties.RecordSet_GeoLocation `yaml:"GeoLocation,omitempty"`
	AliasTarget *properties.RecordSet_AliasTarget `yaml:"AliasTarget,omitempty"`
}

func NewRoute53RecordSet(properties Route53RecordSetProperties, deps ...interface{}) Route53RecordSet {
	return Route53RecordSet{
		Type:       "AWS::Route53::RecordSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRoute53RecordSet(name string, data string) (cf types.ValueMap, err error) {
	var resource Route53RecordSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Route53RecordSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource Route53RecordSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource Route53RecordSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
