package properties

	import "fmt"

type Bucket_TopicConfiguration struct {
	
	
	
	Event interface{} `yaml:"Event"`
	Topic interface{} `yaml:"Topic"`
	Filter *Bucket_NotificationFilter `yaml:"Filter,omitempty"`
}

func (resource Bucket_TopicConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Event == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Event'"))
	}
	if resource.Topic == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Topic'"))
	}
	return errs
}
