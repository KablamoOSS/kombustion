package properties


type DeploymentGroup_RevisionLocation struct {
	
	
	
	RevisionType interface{} `yaml:"RevisionType,omitempty"`
	S3Location *DeploymentGroup_S3Location `yaml:"S3Location,omitempty"`
	GitHubLocation *DeploymentGroup_GitHubLocation `yaml:"GitHubLocation,omitempty"`
}

func (resource DeploymentGroup_RevisionLocation) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
