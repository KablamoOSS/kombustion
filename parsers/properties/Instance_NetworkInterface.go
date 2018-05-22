package properties

	import "fmt"

type Instance_NetworkInterface struct {
	
	
	
	
	
	
	
	
	
	
	
	
	AssociatePublicIpAddress interface{} `yaml:"AssociatePublicIpAddress,omitempty"`
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	DeviceIndex interface{} `yaml:"DeviceIndex"`
	Ipv6AddressCount interface{} `yaml:"Ipv6AddressCount,omitempty"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress,omitempty"`
	SecondaryPrivateIpAddressCount interface{} `yaml:"SecondaryPrivateIpAddressCount,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	GroupSet interface{} `yaml:"GroupSet,omitempty"`
	Ipv6Addresses interface{} `yaml:"Ipv6Addresses,omitempty"`
	PrivateIpAddresses interface{} `yaml:"PrivateIpAddresses,omitempty"`
}

func (resource Instance_NetworkInterface) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.DeviceIndex == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeviceIndex'"))
	}
	return errs
}
