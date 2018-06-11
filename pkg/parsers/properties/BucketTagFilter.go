package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketTagFilter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-tagfilter.html
type BucketTagFilter struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

func (resource BucketTagFilter) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
