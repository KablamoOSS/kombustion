package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceNetworkInterface Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
type InstanceNetworkInterface struct {
	AssociatePublicIpAddress       interface{} `yaml:"AssociatePublicIpAddress,omitempty"`
	DeleteOnTermination            interface{} `yaml:"DeleteOnTermination,omitempty"`
	Description                    interface{} `yaml:"Description,omitempty"`
	DeviceIndex                    interface{} `yaml:"DeviceIndex"`
	Ipv6AddressCount               interface{} `yaml:"Ipv6AddressCount,omitempty"`
	NetworkInterfaceId             interface{} `yaml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress               interface{} `yaml:"PrivateIpAddress,omitempty"`
	SecondaryPrivateIpAddressCount interface{} `yaml:"SecondaryPrivateIpAddressCount,omitempty"`
	SubnetId                       interface{} `yaml:"SubnetId,omitempty"`
	GroupSet                       interface{} `yaml:"GroupSet,omitempty"`
	Ipv6Addresses                  interface{} `yaml:"Ipv6Addresses,omitempty"`
	PrivateIpAddresses             interface{} `yaml:"PrivateIpAddresses,omitempty"`
}

// InstanceNetworkInterface validation
func (resource InstanceNetworkInterface) Validate() []error {
	errors := []error{}

	if resource.DeviceIndex == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DeviceIndex'"))
	}
	return errors
}
