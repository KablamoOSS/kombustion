package properties


type Filter_FindingCriteria struct {
	
	
	Criterion interface{} `yaml:"Criterion,omitempty"`
	ItemType *Filter_Condition `yaml:"ItemType,omitempty"`
}

func (resource Filter_FindingCriteria) Validate() []error {
	errs := []error{}
	
	
	return errs
}
