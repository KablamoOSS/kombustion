package properties

	import "fmt"

type Cluster_PlacementType struct {
	
	AvailabilityZone interface{} `yaml:"AvailabilityZone"`
}

func (resource Cluster_PlacementType) Validate() []error {
	errs := []error{}
	
	if resource.AvailabilityZone == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AvailabilityZone'"))
	}
	return errs
}
