package properties


type PatchBaseline_PatchFilterGroup struct {
	
	PatchFilters interface{} `yaml:"PatchFilters,omitempty"`
}

func (resource PatchBaseline_PatchFilterGroup) Validate() []error {
	errs := []error{}
	
	return errs
}
