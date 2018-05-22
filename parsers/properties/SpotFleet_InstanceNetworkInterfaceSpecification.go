package properties


type SpotFleet_InstanceNetworkInterfaceSpecification struct {
	
	
	
	
	
	
	
	
	
	
	
	AssociatePublicIpAddress interface{} `yaml:"AssociatePublicIpAddress,omitempty"`
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	DeviceIndex interface{} `yaml:"DeviceIndex,omitempty"`
	Ipv6AddressCount interface{} `yaml:"Ipv6AddressCount,omitempty"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId,omitempty"`
	SecondaryPrivateIpAddressCount interface{} `yaml:"SecondaryPrivateIpAddressCount,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	Groups interface{} `yaml:"Groups,omitempty"`
	Ipv6Addresses interface{} `yaml:"Ipv6Addresses,omitempty"`
	PrivateIpAddresses interface{} `yaml:"PrivateIpAddresses,omitempty"`
}

func (resource SpotFleet_InstanceNetworkInterfaceSpecification) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
