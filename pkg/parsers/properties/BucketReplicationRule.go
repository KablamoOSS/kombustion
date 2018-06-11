package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketReplicationRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html
type BucketReplicationRule struct {
	Id                      interface{}                    `yaml:"Id,omitempty"`
	Prefix                  interface{}                    `yaml:"Prefix"`
	Status                  interface{}                    `yaml:"Status"`
	SourceSelectionCriteria *BucketSourceSelectionCriteria `yaml:"SourceSelectionCriteria,omitempty"`
	Destination             *BucketReplicationDestination  `yaml:"Destination"`
}

// BucketReplicationRule validation
func (resource BucketReplicationRule) Validate() []error {
	errs := []error{}

	if resource.Prefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Prefix'"))
	}
	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	if resource.Destination == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Destination'"))
	} else {
		errs = append(errs, resource.Destination.Validate()...)
	}
	return errs
}