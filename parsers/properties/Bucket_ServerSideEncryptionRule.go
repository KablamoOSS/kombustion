package properties


type Bucket_ServerSideEncryptionRule struct {
	
	ServerSideEncryptionByDefault *Bucket_ServerSideEncryptionByDefault `yaml:"ServerSideEncryptionByDefault,omitempty"`
}

func (resource Bucket_ServerSideEncryptionRule) Validate() []error {
	errs := []error{}
	
	return errs
}
