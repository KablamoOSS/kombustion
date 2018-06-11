package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceFleetConfigInstanceFleetProvisioningSpecifications Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html
type InstanceFleetConfigInstanceFleetProvisioningSpecifications struct {
	SpotSpecification *InstanceFleetConfigSpotProvisioningSpecification `yaml:"SpotSpecification"`
}

// InstanceFleetConfigInstanceFleetProvisioningSpecifications validation
func (resource InstanceFleetConfigInstanceFleetProvisioningSpecifications) Validate() []error {
	errs := []error{}

	if resource.SpotSpecification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SpotSpecification'"))
	} else {
		errs = append(errs, resource.SpotSpecification.Validate()...)
	}
	return errs
}