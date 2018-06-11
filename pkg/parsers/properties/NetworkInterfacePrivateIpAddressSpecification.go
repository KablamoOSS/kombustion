package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// NetworkInterfacePrivateIpAddressSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type NetworkInterfacePrivateIpAddressSpecification struct {
	Primary          interface{} `yaml:"Primary"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress"`
}

func (resource NetworkInterfacePrivateIpAddressSpecification) Validate() []error {
	errs := []error{}

	if resource.Primary == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Primary'"))
	}
	if resource.PrivateIpAddress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PrivateIpAddress'"))
	}
	return errs
}
