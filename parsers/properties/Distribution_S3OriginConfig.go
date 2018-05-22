package properties


type Distribution_S3OriginConfig struct {
	
	OriginAccessIdentity interface{} `yaml:"OriginAccessIdentity,omitempty"`
}

func (resource Distribution_S3OriginConfig) Validate() []error {
	errs := []error{}
	
	return errs
}
