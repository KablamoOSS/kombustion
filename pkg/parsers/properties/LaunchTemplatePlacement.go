package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LaunchTemplatePlacement Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-placement.html
type LaunchTemplatePlacement struct {
	Affinity         interface{} `yaml:"Affinity,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	GroupName        interface{} `yaml:"GroupName,omitempty"`
	HostId           interface{} `yaml:"HostId,omitempty"`
	Tenancy          interface{} `yaml:"Tenancy,omitempty"`
}

func (resource LaunchTemplatePlacement) Validate() []error {
	errs := []error{}

	return errs
}
