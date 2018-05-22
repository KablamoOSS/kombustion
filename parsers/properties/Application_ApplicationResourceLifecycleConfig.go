package properties


type Application_ApplicationResourceLifecycleConfig struct {
	
	
	ServiceRole interface{} `yaml:"ServiceRole,omitempty"`
	VersionLifecycleConfig *Application_ApplicationVersionLifecycleConfig `yaml:"VersionLifecycleConfig,omitempty"`
}

func (resource Application_ApplicationResourceLifecycleConfig) Validate() []error {
	errs := []error{}
	
	
	return errs
}
