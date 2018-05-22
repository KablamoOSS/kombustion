package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ServiceDiscoveryService struct {
	Type       string                      `yaml:"Type"`
	Properties ServiceDiscoveryServiceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ServiceDiscoveryServiceProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	HealthCheckConfig *properties.Service_HealthCheckConfig `yaml:"HealthCheckConfig,omitempty"`
	DnsConfig *properties.Service_DnsConfig `yaml:"DnsConfig"`
}

func NewServiceDiscoveryService(properties ServiceDiscoveryServiceProperties, deps ...interface{}) ServiceDiscoveryService {
	return ServiceDiscoveryService{
		Type:       "AWS::ServiceDiscovery::Service",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseServiceDiscoveryService(name string, data string) (cf types.ValueMap, err error) {
	var resource ServiceDiscoveryService
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceDiscoveryService - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ServiceDiscoveryService) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ServiceDiscoveryServiceProperties) Validate() []error {
	errs := []error{}
	if resource.DnsConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DnsConfig'"))
	} else {
		errs = append(errs, resource.DnsConfig.Validate()...)
	}
	return errs
}
