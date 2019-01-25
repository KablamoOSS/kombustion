package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ListenerRuleAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type ListenerRuleAction struct {
	Order                     interface{} `yaml:"Order,omitempty"`
	TargetGroupArn            interface{} `yaml:"TargetGroupArn,omitempty"`
	Type                      interface{} `yaml:"Type"`
	RedirectConfig            interface{} `yaml:"RedirectConfig,omitempty"`
	FixedResponseConfig       interface{} `yaml:"FixedResponseConfig,omitempty"`
	AuthenticateOidcConfig    interface{} `yaml:"AuthenticateOidcConfig,omitempty"`
	AuthenticateCognitoConfig interface{} `yaml:"AuthenticateCognitoConfig,omitempty"`
}

// ListenerRuleAction validation
func (resource ListenerRuleAction) Validate() []error {
	errors := []error{}

	return errors
}
