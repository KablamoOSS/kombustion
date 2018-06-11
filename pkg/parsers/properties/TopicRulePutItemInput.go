package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TopicRulePutItemInput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putiteminput.html
type TopicRulePutItemInput struct {
	TableName interface{} `yaml:"TableName"`
}

func (resource TopicRulePutItemInput) Validate() []error {
	errs := []error{}

	if resource.TableName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableName'"))
	}
	return errs
}
