package properties

	import "fmt"

type Bucket_LambdaConfiguration struct {
	
	
	
	Event interface{} `yaml:"Event"`
	Function interface{} `yaml:"Function"`
	Filter *Bucket_NotificationFilter `yaml:"Filter,omitempty"`
}

func (resource Bucket_LambdaConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Event == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Event'"))
	}
	if resource.Function == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Function'"))
	}
	return errs
}
