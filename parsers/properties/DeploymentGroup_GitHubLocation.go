package properties

	import "fmt"

type DeploymentGroup_GitHubLocation struct {
	
	
	CommitId interface{} `yaml:"CommitId"`
	Repository interface{} `yaml:"Repository"`
}

func (resource DeploymentGroup_GitHubLocation) Validate() []error {
	errs := []error{}
	
	
	if resource.CommitId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CommitId'"))
	}
	if resource.Repository == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Repository'"))
	}
	return errs
}
