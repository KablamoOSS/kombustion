package properties

	import "fmt"

type IdentityPoolRoleAttachment_RoleMapping struct {
	
	
	
	AmbiguousRoleResolution interface{} `yaml:"AmbiguousRoleResolution,omitempty"`
	Type interface{} `yaml:"Type"`
	RulesConfiguration *IdentityPoolRoleAttachment_RulesConfigurationType `yaml:"RulesConfiguration,omitempty"`
}

func (resource IdentityPoolRoleAttachment_RoleMapping) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
