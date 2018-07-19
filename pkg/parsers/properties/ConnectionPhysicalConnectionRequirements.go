package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConnectionPhysicalConnectionRequirements Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html
type ConnectionPhysicalConnectionRequirements struct {
	AvailabilityZone    interface{} `yaml:"AvailabilityZone,omitempty"`
	SubnetId            interface{} `yaml:"SubnetId,omitempty"`
	SecurityGroupIdList interface{} `yaml:"SecurityGroupIdList,omitempty"`
}

// ConnectionPhysicalConnectionRequirements validation
func (resource ConnectionPhysicalConnectionRequirements) Validate() []error {
	errors := []error{}

	return errors
}
