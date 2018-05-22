package properties

	import "fmt"

type Pipeline_InputArtifact struct {
	
	Name interface{} `yaml:"Name"`
}

func (resource Pipeline_InputArtifact) Validate() []error {
	errs := []error{}
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
