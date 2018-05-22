package properties


type Crawler_SchemaChangePolicy struct {
	
	
	DeleteBehavior interface{} `yaml:"DeleteBehavior,omitempty"`
	UpdateBehavior interface{} `yaml:"UpdateBehavior,omitempty"`
}

func (resource Crawler_SchemaChangePolicy) Validate() []error {
	errs := []error{}
	
	
	return errs
}
