package properties


type VPNConnection_VpnTunnelOptionsSpecification struct {
	
	
	PreSharedKey interface{} `yaml:"PreSharedKey,omitempty"`
	TunnelInsideCidr interface{} `yaml:"TunnelInsideCidr,omitempty"`
}

func (resource VPNConnection_VpnTunnelOptionsSpecification) Validate() []error {
	errs := []error{}
	
	
	return errs
}
