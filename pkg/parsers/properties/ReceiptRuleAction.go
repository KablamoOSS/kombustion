package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ReceiptRuleAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html
type ReceiptRuleAction struct {
	WorkmailAction  *ReceiptRuleWorkmailAction  `yaml:"WorkmailAction,omitempty"`
	StopAction      *ReceiptRuleStopAction      `yaml:"StopAction,omitempty"`
	SNSAction       *ReceiptRuleSNSAction       `yaml:"SNSAction,omitempty"`
	S3Action        *ReceiptRuleS3Action        `yaml:"S3Action,omitempty"`
	LambdaAction    *ReceiptRuleLambdaAction    `yaml:"LambdaAction,omitempty"`
	BounceAction    *ReceiptRuleBounceAction    `yaml:"BounceAction,omitempty"`
	AddHeaderAction *ReceiptRuleAddHeaderAction `yaml:"AddHeaderAction,omitempty"`
}

func (resource ReceiptRuleAction) Validate() []error {
	errs := []error{}

	return errs
}
