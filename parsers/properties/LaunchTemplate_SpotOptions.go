package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type LaunchTemplate_SpotOptions struct {
	InstanceInterruptionBehavior interface{} `yaml:"InstanceInterruptionBehavior,omitempty"`
	MaxPrice                     interface{} `yaml:"MaxPrice,omitempty"`
	SpotInstanceType             interface{} `yaml:"SpotInstanceType,omitempty"`
}

func (resource LaunchTemplate_SpotOptions) Validate() []error {
	errs := []error{}

	return errs
}
