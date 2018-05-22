package properties

	import "fmt"

type ListenerRule_Action struct {
	
	
	TargetGroupArn interface{} `yaml:"TargetGroupArn"`
	Type interface{} `yaml:"Type"`
}

func (resource ListenerRule_Action) Validate() []error {
	errs := []error{}
	
	
	if resource.TargetGroupArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetGroupArn'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
