package properties

	import "fmt"

type Cluster_ScalingTrigger struct {
	
	CloudWatchAlarmDefinition *Cluster_CloudWatchAlarmDefinition `yaml:"CloudWatchAlarmDefinition"`
}

func (resource Cluster_ScalingTrigger) Validate() []error {
	errs := []error{}
	
	if resource.CloudWatchAlarmDefinition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CloudWatchAlarmDefinition'"))
	} else {
		errs = append(errs, resource.CloudWatchAlarmDefinition.Validate()...)
	}
	return errs
}
