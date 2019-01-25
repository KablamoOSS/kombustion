package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicRuleStepFunctionsAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html
type TopicRuleStepFunctionsAction struct {
	ExecutionNamePrefix interface{} `yaml:"ExecutionNamePrefix,omitempty"`
	RoleArn             interface{} `yaml:"RoleArn"`
	StateMachineName    interface{} `yaml:"StateMachineName"`
}

// TopicRuleStepFunctionsAction validation
func (resource TopicRuleStepFunctionsAction) Validate() []error {
	errors := []error{}

	return errors
}