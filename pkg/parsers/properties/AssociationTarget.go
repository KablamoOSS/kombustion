package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// AssociationTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
type AssociationTarget struct {
	Key    interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values"`
}

func (resource AssociationTarget) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Values == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Values'"))
	}
	return errs
}
