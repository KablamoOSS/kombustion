package properties


type SpotFleet_LaunchTemplateConfig struct {
	
	
	Overrides interface{} `yaml:"Overrides,omitempty"`
	LaunchTemplateSpecification *SpotFleet_FleetLaunchTemplateSpecification `yaml:"LaunchTemplateSpecification,omitempty"`
}

func (resource SpotFleet_LaunchTemplateConfig) Validate() []error {
	errs := []error{}
	
	
	return errs
}
