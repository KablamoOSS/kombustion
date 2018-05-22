package properties

	import "fmt"

type InstanceGroupConfig_CloudWatchAlarmDefinition struct {
	
	
	
	
	
	
	
	
	
	ComparisonOperator interface{} `yaml:"ComparisonOperator"`
	EvaluationPeriods interface{} `yaml:"EvaluationPeriods,omitempty"`
	MetricName interface{} `yaml:"MetricName"`
	Namespace interface{} `yaml:"Namespace,omitempty"`
	Period interface{} `yaml:"Period"`
	Statistic interface{} `yaml:"Statistic,omitempty"`
	Threshold interface{} `yaml:"Threshold"`
	Unit interface{} `yaml:"Unit,omitempty"`
	Dimensions interface{} `yaml:"Dimensions,omitempty"`
}

func (resource InstanceGroupConfig_CloudWatchAlarmDefinition) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	if resource.ComparisonOperator == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComparisonOperator'"))
	}
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Period == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Period'"))
	}
	if resource.Threshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Threshold'"))
	}
	return errs
}
