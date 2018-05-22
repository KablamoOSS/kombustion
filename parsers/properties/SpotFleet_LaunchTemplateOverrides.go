package properties


type SpotFleet_LaunchTemplateOverrides struct {
	
	
	
	
	
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	InstanceType interface{} `yaml:"InstanceType,omitempty"`
	SpotPrice interface{} `yaml:"SpotPrice,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	WeightedCapacity interface{} `yaml:"WeightedCapacity,omitempty"`
}

func (resource SpotFleet_LaunchTemplateOverrides) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
