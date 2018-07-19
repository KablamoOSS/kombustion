package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketTopicConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html
type BucketTopicConfiguration struct {
	Event  interface{}               `yaml:"Event"`
	Topic  interface{}               `yaml:"Topic"`
	Filter *BucketNotificationFilter `yaml:"Filter,omitempty"`
}

// BucketTopicConfiguration validation
func (resource BucketTopicConfiguration) Validate() []error {
	errors := []error{}

	if resource.Event == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Event'"))
	}
	if resource.Topic == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Topic'"))
	}
	return errors
}
