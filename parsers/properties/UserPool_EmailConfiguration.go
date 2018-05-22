package properties


type UserPool_EmailConfiguration struct {
	
	
	ReplyToEmailAddress interface{} `yaml:"ReplyToEmailAddress,omitempty"`
	SourceArn interface{} `yaml:"SourceArn,omitempty"`
}

func (resource UserPool_EmailConfiguration) Validate() []error {
	errs := []error{}
	
	
	return errs
}
