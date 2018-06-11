package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SpotFleetSpotPlacement Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html
type SpotFleetSpotPlacement struct {
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	GroupName        interface{} `yaml:"GroupName,omitempty"`
}

func (resource SpotFleetSpotPlacement) Validate() []error {
	errs := []error{}

	return errs
}
