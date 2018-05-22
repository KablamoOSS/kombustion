package properties

	import "fmt"

type ReceiptRule_BounceAction struct {
	
	
	
	
	
	Message interface{} `yaml:"Message"`
	Sender interface{} `yaml:"Sender"`
	SmtpReplyCode interface{} `yaml:"SmtpReplyCode"`
	StatusCode interface{} `yaml:"StatusCode,omitempty"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

func (resource ReceiptRule_BounceAction) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.Message == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Message'"))
	}
	if resource.Sender == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Sender'"))
	}
	if resource.SmtpReplyCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SmtpReplyCode'"))
	}
	return errs
}
