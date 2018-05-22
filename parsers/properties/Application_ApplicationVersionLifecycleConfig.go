package properties


type Application_ApplicationVersionLifecycleConfig struct {
	
	
	MaxCountRule *Application_MaxCountRule `yaml:"MaxCountRule,omitempty"`
	MaxAgeRule *Application_MaxAgeRule `yaml:"MaxAgeRule,omitempty"`
}

func (resource Application_ApplicationVersionLifecycleConfig) Validate() []error {
	errs := []error{}
	
	
	return errs
}
