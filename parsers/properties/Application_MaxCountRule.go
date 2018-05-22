package properties


type Application_MaxCountRule struct {
	
	
	
	DeleteSourceFromS3 interface{} `yaml:"DeleteSourceFromS3,omitempty"`
	Enabled interface{} `yaml:"Enabled,omitempty"`
	MaxCount interface{} `yaml:"MaxCount,omitempty"`
}

func (resource Application_MaxCountRule) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
