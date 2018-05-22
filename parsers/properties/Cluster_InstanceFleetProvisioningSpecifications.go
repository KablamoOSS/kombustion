package properties

	import "fmt"

type Cluster_InstanceFleetProvisioningSpecifications struct {
	
	SpotSpecification *Cluster_SpotProvisioningSpecification `yaml:"SpotSpecification"`
}

func (resource Cluster_InstanceFleetProvisioningSpecifications) Validate() []error {
	errs := []error{}
	
	if resource.SpotSpecification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SpotSpecification'"))
	} else {
		errs = append(errs, resource.SpotSpecification.Validate()...)
	}
	return errs
}
