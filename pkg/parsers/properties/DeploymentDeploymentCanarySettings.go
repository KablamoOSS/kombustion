package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentDeploymentCanarySettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html
type DeploymentDeploymentCanarySettings struct {
	PercentTraffic         interface{} `yaml:"PercentTraffic,omitempty"`
	UseStageCache          interface{} `yaml:"UseStageCache,omitempty"`
	StageVariableOverrides interface{} `yaml:"StageVariableOverrides,omitempty"`
}

// DeploymentDeploymentCanarySettings validation
func (resource DeploymentDeploymentCanarySettings) Validate() []error {
	errors := []error{}

	return errors
}