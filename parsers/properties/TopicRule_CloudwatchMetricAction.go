package properties

	import "fmt"

type TopicRule_CloudwatchMetricAction struct {
	
	
	
	
	
	
	MetricName interface{} `yaml:"MetricName"`
	MetricNamespace interface{} `yaml:"MetricNamespace"`
	MetricTimestamp interface{} `yaml:"MetricTimestamp,omitempty"`
	MetricUnit interface{} `yaml:"MetricUnit"`
	MetricValue interface{} `yaml:"MetricValue"`
	RoleArn interface{} `yaml:"RoleArn"`
}

func (resource TopicRule_CloudwatchMetricAction) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.MetricNamespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricNamespace'"))
	}
	if resource.MetricUnit == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricUnit'"))
	}
	if resource.MetricValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricValue'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
