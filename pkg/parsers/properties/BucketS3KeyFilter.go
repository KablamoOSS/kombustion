package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketS3KeyFilter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
type BucketS3KeyFilter struct {
	Rules interface{} `yaml:"Rules"`
}

func (resource BucketS3KeyFilter) Validate() []error {
	errs := []error{}

	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
