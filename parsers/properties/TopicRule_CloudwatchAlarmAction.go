package properties

	import "fmt"

type TopicRule_CloudwatchAlarmAction struct {
	
	
	
	
	AlarmName interface{} `yaml:"AlarmName"`
	RoleArn interface{} `yaml:"RoleArn"`
	StateReason interface{} `yaml:"StateReason"`
	StateValue interface{} `yaml:"StateValue"`
}

func (resource TopicRule_CloudwatchAlarmAction) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.AlarmName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AlarmName'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.StateReason == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StateReason'"))
	}
	if resource.StateValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StateValue'"))
	}
	return errs
}
