package properties


type LaunchTemplate_InstanceMarketOptions struct {
	
	
	MarketType interface{} `yaml:"MarketType,omitempty"`
	SpotOptions *LaunchTemplate_SpotOptions `yaml:"SpotOptions,omitempty"`
}

func (resource LaunchTemplate_InstanceMarketOptions) Validate() []error {
	errs := []error{}
	
	
	return errs
}
