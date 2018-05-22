package properties


type ReceiptRule_Action struct {
	
	
	
	
	
	
	
	WorkmailAction *ReceiptRule_WorkmailAction `yaml:"WorkmailAction,omitempty"`
	StopAction *ReceiptRule_StopAction `yaml:"StopAction,omitempty"`
	SNSAction *ReceiptRule_SNSAction `yaml:"SNSAction,omitempty"`
	S3Action *ReceiptRule_S3Action `yaml:"S3Action,omitempty"`
	LambdaAction *ReceiptRule_LambdaAction `yaml:"LambdaAction,omitempty"`
	BounceAction *ReceiptRule_BounceAction `yaml:"BounceAction,omitempty"`
	AddHeaderAction *ReceiptRule_AddHeaderAction `yaml:"AddHeaderAction,omitempty"`
}

func (resource ReceiptRule_Action) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	return errs
}
