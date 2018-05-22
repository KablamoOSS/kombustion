package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ServiceDiscoveryInstance struct {
	Type       string                      `yaml:"Type"`
	Properties ServiceDiscoveryInstanceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ServiceDiscoveryInstanceProperties struct {
	InstanceAttributes interface{} `yaml:"InstanceAttributes"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
	ServiceId interface{} `yaml:"ServiceId"`
}

func NewServiceDiscoveryInstance(properties ServiceDiscoveryInstanceProperties, deps ...interface{}) ServiceDiscoveryInstance {
	return ServiceDiscoveryInstance{
		Type:       "AWS::ServiceDiscovery::Instance",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseServiceDiscoveryInstance(name string, data string) (cf types.ValueMap, err error) {
	var resource ServiceDiscoveryInstance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceDiscoveryInstance - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ServiceDiscoveryInstance) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ServiceDiscoveryInstanceProperties) Validate() []error {
	errs := []error{}
	if resource.InstanceAttributes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceAttributes'"))
	}
	if resource.ServiceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceId'"))
	}
	return errs
}
