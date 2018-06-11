package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeploymentGroupGitHubLocation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html
type DeploymentGroupGitHubLocation struct {
	CommitId   interface{} `yaml:"CommitId"`
	Repository interface{} `yaml:"Repository"`
}

func (resource DeploymentGroupGitHubLocation) Validate() []error {
	errs := []error{}

	if resource.CommitId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CommitId'"))
	}
	if resource.Repository == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Repository'"))
	}
	return errs
}
