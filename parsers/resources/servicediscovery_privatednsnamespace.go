package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ServiceDiscoveryPrivateDnsNamespace struct {
	Type       string                      `yaml:"Type"`
	Properties ServiceDiscoveryPrivateDnsNamespaceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ServiceDiscoveryPrivateDnsNamespaceProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	Vpc interface{} `yaml:"Vpc"`
}

func NewServiceDiscoveryPrivateDnsNamespace(properties ServiceDiscoveryPrivateDnsNamespaceProperties, deps ...interface{}) ServiceDiscoveryPrivateDnsNamespace {
	return ServiceDiscoveryPrivateDnsNamespace{
		Type:       "AWS::ServiceDiscovery::PrivateDnsNamespace",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseServiceDiscoveryPrivateDnsNamespace(name string, data string) (cf types.ValueMap, err error) {
	var resource ServiceDiscoveryPrivateDnsNamespace
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceDiscoveryPrivateDnsNamespace - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ServiceDiscoveryPrivateDnsNamespace) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ServiceDiscoveryPrivateDnsNamespaceProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Vpc == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Vpc'"))
	}
	return errs
}
