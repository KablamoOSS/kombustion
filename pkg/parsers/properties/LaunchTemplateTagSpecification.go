package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LaunchTemplateTagSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html
type LaunchTemplateTagSpecification struct {
	ResourceType interface{} `yaml:"ResourceType,omitempty"`
	Tags         interface{} `yaml:"Tags,omitempty"`
}

func (resource LaunchTemplateTagSpecification) Validate() []error {
	errs := []error{}

	return errs
}
