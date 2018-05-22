package properties

	import "fmt"

type AutoScalingGroup_TagProperty struct {
	
	
	
	Key interface{} `yaml:"Key"`
	PropagateAtLaunch interface{} `yaml:"PropagateAtLaunch"`
	Value interface{} `yaml:"Value"`
}

func (resource AutoScalingGroup_TagProperty) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.PropagateAtLaunch == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PropagateAtLaunch'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
