package properties


type Bucket_StorageClassAnalysis struct {
	
	DataExport *Bucket_DataExport `yaml:"DataExport,omitempty"`
}

func (resource Bucket_StorageClassAnalysis) Validate() []error {
	errs := []error{}
	
	return errs
}
