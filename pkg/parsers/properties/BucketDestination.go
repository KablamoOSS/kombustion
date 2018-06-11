package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketDestination Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html
type BucketDestination struct {
	BucketAccountId interface{} `yaml:"BucketAccountId,omitempty"`
	BucketArn       interface{} `yaml:"BucketArn"`
	Format          interface{} `yaml:"Format"`
	Prefix          interface{} `yaml:"Prefix,omitempty"`
}

func (resource BucketDestination) Validate() []error {
	errs := []error{}

	if resource.BucketArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketArn'"))
	}
	if resource.Format == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Format'"))
	}
	return errs
}
