package properties


type Layer_LoadBasedAutoScaling struct {
	
	
	
	Enable interface{} `yaml:"Enable,omitempty"`
	DownScaling *Layer_AutoScalingThresholds `yaml:"DownScaling,omitempty"`
	UpScaling *Layer_AutoScalingThresholds `yaml:"UpScaling,omitempty"`
}

func (resource Layer_LoadBasedAutoScaling) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
