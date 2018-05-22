package properties

	import "fmt"

type EnvironmentEC2_Repository struct {
	
	
	PathComponent interface{} `yaml:"PathComponent"`
	RepositoryUrl interface{} `yaml:"RepositoryUrl"`
}

func (resource EnvironmentEC2_Repository) Validate() []error {
	errs := []error{}
	
	
	if resource.PathComponent == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PathComponent'"))
	}
	if resource.RepositoryUrl == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RepositoryUrl'"))
	}
	return errs
}
