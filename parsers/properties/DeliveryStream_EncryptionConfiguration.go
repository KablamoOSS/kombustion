package properties


type DeliveryStream_EncryptionConfiguration struct {
	
	
	NoEncryptionConfig interface{} `yaml:"NoEncryptionConfig,omitempty"`
	KMSEncryptionConfig *DeliveryStream_KMSEncryptionConfig `yaml:"KMSEncryptionConfig,omitempty"`
}

func (resource DeliveryStream_EncryptionConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
