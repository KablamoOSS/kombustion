package properties


type LaunchTemplate_Placement struct {
	
	
	
	
	
	Affinity interface{} `yaml:"Affinity,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	GroupName interface{} `yaml:"GroupName,omitempty"`
	HostId interface{} `yaml:"HostId,omitempty"`
	Tenancy interface{} `yaml:"Tenancy,omitempty"`
}

func (resource LaunchTemplate_Placement) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
