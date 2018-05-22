package properties

	import "fmt"

type TopicRule_SqsAction struct {
	
	
	
	QueueUrl interface{} `yaml:"QueueUrl"`
	RoleArn interface{} `yaml:"RoleArn"`
	UseBase64 interface{} `yaml:"UseBase64,omitempty"`
}

func (resource TopicRule_SqsAction) Validate() []error {
	errs := []error{}
	
	
	
	if resource.QueueUrl == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'QueueUrl'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
