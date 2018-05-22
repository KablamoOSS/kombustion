package properties

	import "fmt"

type IdentityPoolRoleAttachment_RulesConfigurationType struct {
	
	Rules interface{} `yaml:"Rules"`
}

func (resource IdentityPoolRoleAttachment_RulesConfigurationType) Validate() []error {
	errs := []error{}
	
	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
