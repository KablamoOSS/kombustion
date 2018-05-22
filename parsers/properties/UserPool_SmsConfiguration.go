package properties


type UserPool_SmsConfiguration struct {
	
	
	ExternalId interface{} `yaml:"ExternalId,omitempty"`
	SnsCallerArn interface{} `yaml:"SnsCallerArn,omitempty"`
}

func (resource UserPool_SmsConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
