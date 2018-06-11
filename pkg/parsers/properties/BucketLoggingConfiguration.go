package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// BucketLoggingConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html
type BucketLoggingConfiguration struct {
	DestinationBucketName interface{} `yaml:"DestinationBucketName,omitempty"`
	LogFilePrefix         interface{} `yaml:"LogFilePrefix,omitempty"`
}

func (resource BucketLoggingConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
