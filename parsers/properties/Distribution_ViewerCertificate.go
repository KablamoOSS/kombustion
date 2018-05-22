package properties


type Distribution_ViewerCertificate struct {
	
	
	
	
	
	AcmCertificateArn interface{} `yaml:"AcmCertificateArn,omitempty"`
	CloudFrontDefaultCertificate interface{} `yaml:"CloudFrontDefaultCertificate,omitempty"`
	IamCertificateId interface{} `yaml:"IamCertificateId,omitempty"`
	MinimumProtocolVersion interface{} `yaml:"MinimumProtocolVersion,omitempty"`
	SslSupportMethod interface{} `yaml:"SslSupportMethod,omitempty"`
}

func (resource Distribution_ViewerCertificate) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
