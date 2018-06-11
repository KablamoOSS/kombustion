package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentGroupTriggerConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html
type DeploymentGroupTriggerConfig struct {
	TriggerName      interface{} `yaml:"TriggerName,omitempty"`
	TriggerTargetArn interface{} `yaml:"TriggerTargetArn,omitempty"`
	TriggerEvents    interface{} `yaml:"TriggerEvents,omitempty"`
}

func (resource DeploymentGroupTriggerConfig) Validate() []error {
	errs := []error{}

	return errs
}
