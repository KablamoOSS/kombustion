package properties

	import "fmt"

type Pipeline_ActionDeclaration struct {
	
	
	
	
	
	
	
	Configuration interface{} `yaml:"Configuration,omitempty"`
	Name interface{} `yaml:"Name"`
	RoleArn interface{} `yaml:"RoleArn,omitempty"`
	RunOrder interface{} `yaml:"RunOrder,omitempty"`
	InputArtifacts interface{} `yaml:"InputArtifacts,omitempty"`
	OutputArtifacts interface{} `yaml:"OutputArtifacts,omitempty"`
	ActionTypeId *Pipeline_ActionTypeId `yaml:"ActionTypeId"`
}

func (resource Pipeline_ActionDeclaration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ActionTypeId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ActionTypeId'"))
	} else {
		errs = append(errs, resource.ActionTypeId.Validate()...)
	}
	return errs
}
