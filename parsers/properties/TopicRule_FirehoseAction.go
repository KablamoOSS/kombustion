package properties

	import "fmt"

type TopicRule_FirehoseAction struct {
	
	
	
	DeliveryStreamName interface{} `yaml:"DeliveryStreamName"`
	RoleArn interface{} `yaml:"RoleArn"`
	Separator interface{} `yaml:"Separator,omitempty"`
}

func (resource TopicRule_FirehoseAction) Validate() []error {
	errs := []error{}
	
	
	
	if resource.DeliveryStreamName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeliveryStreamName'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
