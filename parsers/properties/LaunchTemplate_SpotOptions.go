package properties


type LaunchTemplate_SpotOptions struct {
	
	
	
	InstanceInterruptionBehavior interface{} `yaml:"InstanceInterruptionBehavior,omitempty"`
	MaxPrice interface{} `yaml:"MaxPrice,omitempty"`
	SpotInstanceType interface{} `yaml:"SpotInstanceType,omitempty"`
}

func (resource LaunchTemplate_SpotOptions) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
