package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TopicRuleRepublishAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html
type TopicRuleRepublishAction struct {
	RoleArn interface{} `yaml:"RoleArn"`
	Topic   interface{} `yaml:"Topic"`
}

func (resource TopicRuleRepublishAction) Validate() []error {
	errs := []error{}

	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.Topic == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Topic'"))
	}
	return errs
}
