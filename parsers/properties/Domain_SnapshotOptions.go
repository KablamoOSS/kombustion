package properties


type Domain_SnapshotOptions struct {
	
	AutomatedSnapshotStartHour interface{} `yaml:"AutomatedSnapshotStartHour,omitempty"`
}

func (resource Domain_SnapshotOptions) Validate() []error {
	errs := []error{}
	
	return errs
}
