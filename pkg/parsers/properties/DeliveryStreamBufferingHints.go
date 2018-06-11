package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamBufferingHints Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-bufferinghints.html
type DeliveryStreamBufferingHints struct {
	IntervalInSeconds interface{} `yaml:"IntervalInSeconds"`
	SizeInMBs         interface{} `yaml:"SizeInMBs"`
}

func (resource DeliveryStreamBufferingHints) Validate() []error {
	errs := []error{}

	if resource.IntervalInSeconds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IntervalInSeconds'"))
	}
	if resource.SizeInMBs == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SizeInMBs'"))
	}
	return errs
}
