package properties


type Crawler_Targets struct {
	
	
	JdbcTargets interface{} `yaml:"JdbcTargets,omitempty"`
	S3Targets interface{} `yaml:"S3Targets,omitempty"`
}

func (resource Crawler_Targets) Validate() []error {
	errs := []error{}
	
	
	return errs
}
