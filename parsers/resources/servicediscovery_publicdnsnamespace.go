package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ServiceDiscoveryPublicDnsNamespace struct {
	Type       string                      `yaml:"Type"`
	Properties ServiceDiscoveryPublicDnsNamespaceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ServiceDiscoveryPublicDnsNamespaceProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
}

func NewServiceDiscoveryPublicDnsNamespace(properties ServiceDiscoveryPublicDnsNamespaceProperties, deps ...interface{}) ServiceDiscoveryPublicDnsNamespace {
	return ServiceDiscoveryPublicDnsNamespace{
		Type:       "AWS::ServiceDiscovery::PublicDnsNamespace",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseServiceDiscoveryPublicDnsNamespace(name string, data string) (cf types.ValueMap, err error) {
	var resource ServiceDiscoveryPublicDnsNamespace
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceDiscoveryPublicDnsNamespace - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ServiceDiscoveryPublicDnsNamespace) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ServiceDiscoveryPublicDnsNamespaceProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
