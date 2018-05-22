package properties

	import "fmt"

type HostedZone_QueryLoggingConfig struct {
	
	CloudWatchLogsLogGroupArn interface{} `yaml:"CloudWatchLogsLogGroupArn"`
}

func (resource HostedZone_QueryLoggingConfig) Validate() []error {
	errs := []error{}
	
	if resource.CloudWatchLogsLogGroupArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CloudWatchLogsLogGroupArn'"))
	}
	return errs
}
