package properties

	import "fmt"

type ScalingPlan_TagFilter struct {
	
	
	Key interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values,omitempty"`
}

func (resource ScalingPlan_TagFilter) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}
