package properties

	import "fmt"

type Bucket_QueueConfiguration struct {
	
	
	
	Event interface{} `yaml:"Event"`
	Queue interface{} `yaml:"Queue"`
	Filter *Bucket_NotificationFilter `yaml:"Filter,omitempty"`
}

func (resource Bucket_QueueConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Event == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Event'"))
	}
	if resource.Queue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Queue'"))
	}
	return errs
}
