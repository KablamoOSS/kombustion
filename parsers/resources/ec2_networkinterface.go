package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EC2NetworkInterface struct {
	Type       string                      `yaml:"Type"`
	Properties EC2NetworkInterfaceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2NetworkInterfaceProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	InterfaceType interface{} `yaml:"InterfaceType,omitempty"`
	Ipv6AddressCount interface{} `yaml:"Ipv6AddressCount,omitempty"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress,omitempty"`
	SecondaryPrivateIpAddressCount interface{} `yaml:"SecondaryPrivateIpAddressCount,omitempty"`
	SourceDestCheck interface{} `yaml:"SourceDestCheck,omitempty"`
	SubnetId interface{} `yaml:"SubnetId"`
	GroupSet interface{} `yaml:"GroupSet,omitempty"`
	PrivateIpAddresses interface{} `yaml:"PrivateIpAddresses,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	Ipv6Addresses *properties.NetworkInterface_InstanceIpv6Address `yaml:"Ipv6Addresses,omitempty"`
}

func NewEC2NetworkInterface(properties EC2NetworkInterfaceProperties, deps ...interface{}) EC2NetworkInterface {
	return EC2NetworkInterface{
		Type:       "AWS::EC2::NetworkInterface",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2NetworkInterface(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2NetworkInterface
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2NetworkInterface - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2NetworkInterface) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2NetworkInterfaceProperties) Validate() []error {
	errs := []error{}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errs
}
