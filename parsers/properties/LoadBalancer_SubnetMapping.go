package properties

	import "fmt"

type LoadBalancer_SubnetMapping struct {
	
	
	AllocationId interface{} `yaml:"AllocationId"`
	SubnetId interface{} `yaml:"SubnetId"`
}

func (resource LoadBalancer_SubnetMapping) Validate() []error {
	errs := []error{}
	
	
	if resource.AllocationId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AllocationId'"))
	}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errs
}
