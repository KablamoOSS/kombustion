package properties


type Service_NetworkConfiguration struct {
	
	AwsvpcConfiguration *Service_AwsVpcConfiguration `yaml:"AwsvpcConfiguration,omitempty"`
}

func (resource Service_NetworkConfiguration) Validate() []error {
	errs := []error{}
	
	return errs
}
