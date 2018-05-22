package properties

	import "fmt"

type InstanceFleetConfig_InstanceFleetProvisioningSpecifications struct {
	
	SpotSpecification *InstanceFleetConfig_SpotProvisioningSpecification `yaml:"SpotSpecification"`
}

func (resource InstanceFleetConfig_InstanceFleetProvisioningSpecifications) Validate() []error {
	errs := []error{}
	
	if resource.SpotSpecification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SpotSpecification'"))
	} else {
		errs = append(errs, resource.SpotSpecification.Validate()...)
	}
	return errs
}
