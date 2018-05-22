package properties

	import "fmt"

type ScalableTarget_ScheduledAction struct {
	
	
	
	
	
	EndTime interface{} `yaml:"EndTime,omitempty"`
	Schedule interface{} `yaml:"Schedule"`
	ScheduledActionName interface{} `yaml:"ScheduledActionName"`
	StartTime interface{} `yaml:"StartTime,omitempty"`
	ScalableTargetAction *ScalableTarget_ScalableTargetAction `yaml:"ScalableTargetAction,omitempty"`
}

func (resource ScalableTarget_ScheduledAction) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.Schedule == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Schedule'"))
	}
	if resource.ScheduledActionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScheduledActionName'"))
	}
	return errs
}
