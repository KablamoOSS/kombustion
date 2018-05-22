package properties

	import "fmt"

type Listener_Action struct {
	
	
	TargetGroupArn interface{} `yaml:"TargetGroupArn"`
	Type interface{} `yaml:"Type"`
}

func (resource Listener_Action) Validate() []error {
	errs := []error{}
	
	
	if resource.TargetGroupArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetGroupArn'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
