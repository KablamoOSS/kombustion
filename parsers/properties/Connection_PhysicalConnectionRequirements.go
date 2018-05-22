package properties


type Connection_PhysicalConnectionRequirements struct {
	
	
	
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	SecurityGroupIdList interface{} `yaml:"SecurityGroupIdList,omitempty"`
}

func (resource Connection_PhysicalConnectionRequirements) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
