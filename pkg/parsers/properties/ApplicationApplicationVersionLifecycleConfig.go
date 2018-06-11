package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationApplicationVersionLifecycleConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html
type ApplicationApplicationVersionLifecycleConfig struct {
	MaxCountRule *ApplicationMaxCountRule `yaml:"MaxCountRule,omitempty"`
	MaxAgeRule   *ApplicationMaxAgeRule   `yaml:"MaxAgeRule,omitempty"`
}

func (resource ApplicationApplicationVersionLifecycleConfig) Validate() []error {
	errs := []error{}

	return errs
}
