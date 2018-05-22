package properties


type Application_MaxAgeRule struct {
	
	
	
	DeleteSourceFromS3 interface{} `yaml:"DeleteSourceFromS3,omitempty"`
	Enabled interface{} `yaml:"Enabled,omitempty"`
	MaxAgeInDays interface{} `yaml:"MaxAgeInDays,omitempty"`
}

func (resource Application_MaxAgeRule) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
