package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceFleetConfigSpotProvisioningSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotprovisioningspecification.html
type InstanceFleetConfigSpotProvisioningSpecification struct {
	BlockDurationMinutes   interface{} `yaml:"BlockDurationMinutes,omitempty"`
	TimeoutAction          interface{} `yaml:"TimeoutAction"`
	TimeoutDurationMinutes interface{} `yaml:"TimeoutDurationMinutes"`
}

func (resource InstanceFleetConfigSpotProvisioningSpecification) Validate() []error {
	errs := []error{}

	if resource.TimeoutAction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TimeoutAction'"))
	}
	if resource.TimeoutDurationMinutes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TimeoutDurationMinutes'"))
	}
	return errs
}
