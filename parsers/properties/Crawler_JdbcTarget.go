package properties


type Crawler_JdbcTarget struct {
	
	
	
	ConnectionName interface{} `yaml:"ConnectionName,omitempty"`
	Path interface{} `yaml:"Path,omitempty"`
	Exclusions interface{} `yaml:"Exclusions,omitempty"`
}

func (resource Crawler_JdbcTarget) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
