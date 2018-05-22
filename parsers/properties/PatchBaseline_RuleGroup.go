package properties


type PatchBaseline_RuleGroup struct {
	
	PatchRules interface{} `yaml:"PatchRules,omitempty"`
}

func (resource PatchBaseline_RuleGroup) Validate() []error {
	errs := []error{}
	
	return errs
}
