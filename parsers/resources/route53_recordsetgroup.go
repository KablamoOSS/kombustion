package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type Route53RecordSetGroup struct {
	Type       string                      `yaml:"Type"`
	Properties Route53RecordSetGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type Route53RecordSetGroupProperties struct {
	Comment interface{} `yaml:"Comment,omitempty"`
	HostedZoneId interface{} `yaml:"HostedZoneId,omitempty"`
	HostedZoneName interface{} `yaml:"HostedZoneName,omitempty"`
	RecordSets interface{} `yaml:"RecordSets,omitempty"`
}

func NewRoute53RecordSetGroup(properties Route53RecordSetGroupProperties, deps ...interface{}) Route53RecordSetGroup {
	return Route53RecordSetGroup{
		Type:       "AWS::Route53::RecordSetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRoute53RecordSetGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource Route53RecordSetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Route53RecordSetGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource Route53RecordSetGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource Route53RecordSetGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
