package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterLoggingProperties Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type ClusterLoggingProperties struct {
	BucketName  interface{} `yaml:"BucketName"`
	S3KeyPrefix interface{} `yaml:"S3KeyPrefix,omitempty"`
}

func (resource ClusterLoggingProperties) Validate() []error {
	errs := []error{}

	if resource.BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketName'"))
	}
	return errs
}
