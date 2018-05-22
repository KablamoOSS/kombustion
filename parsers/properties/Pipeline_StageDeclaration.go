package properties

	import "fmt"

type Pipeline_StageDeclaration struct {
	
	
	
	Name interface{} `yaml:"Name"`
	Actions interface{} `yaml:"Actions"`
	Blockers interface{} `yaml:"Blockers,omitempty"`
}

func (resource Pipeline_StageDeclaration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Actions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Actions'"))
	}
	return errs
}
