package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentGroupAutoRollbackConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-autorollbackconfiguration.html
type DeploymentGroupAutoRollbackConfiguration struct {
	Enabled interface{} `yaml:"Enabled,omitempty"`
	Events  interface{} `yaml:"Events,omitempty"`
}

func (resource DeploymentGroupAutoRollbackConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
