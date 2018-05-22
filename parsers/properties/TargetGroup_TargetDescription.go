package properties

	import "fmt"

type TargetGroup_TargetDescription struct {
	
	
	
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	Id interface{} `yaml:"Id"`
	Port interface{} `yaml:"Port,omitempty"`
}

func (resource TargetGroup_TargetDescription) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	return errs
}
