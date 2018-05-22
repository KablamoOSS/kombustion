package properties

	import "fmt"

type InstanceGroupConfig_ScalingTrigger struct {
	
	CloudWatchAlarmDefinition *InstanceGroupConfig_CloudWatchAlarmDefinition `yaml:"CloudWatchAlarmDefinition"`
}

func (resource InstanceGroupConfig_ScalingTrigger) Validate() []error {
	errs := []error{}
	
	if resource.CloudWatchAlarmDefinition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CloudWatchAlarmDefinition'"))
	} else {
		errs = append(errs, resource.CloudWatchAlarmDefinition.Validate()...)
	}
	return errs
}
