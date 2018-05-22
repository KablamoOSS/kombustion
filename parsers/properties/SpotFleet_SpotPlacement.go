package properties


type SpotFleet_SpotPlacement struct {
	
	
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	GroupName interface{} `yaml:"GroupName,omitempty"`
}

func (resource SpotFleet_SpotPlacement) Validate() []error {
	errs := []error{}
	
	
	return errs
}
