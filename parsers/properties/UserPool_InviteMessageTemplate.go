package properties


type UserPool_InviteMessageTemplate struct {
	
	
	
	EmailMessage interface{} `yaml:"EmailMessage,omitempty"`
	EmailSubject interface{} `yaml:"EmailSubject,omitempty"`
	SMSMessage interface{} `yaml:"SMSMessage,omitempty"`
}

func (resource UserPool_InviteMessageTemplate) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
