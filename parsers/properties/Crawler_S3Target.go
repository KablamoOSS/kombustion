package properties


type Crawler_S3Target struct {
	
	
	Path interface{} `yaml:"Path,omitempty"`
	Exclusions interface{} `yaml:"Exclusions,omitempty"`
}

func (resource Crawler_S3Target) Validate() []error {
	errs := []error{}
	
	
	return errs
}
