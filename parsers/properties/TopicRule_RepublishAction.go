package properties

	import "fmt"

type TopicRule_RepublishAction struct {
	
	
	RoleArn interface{} `yaml:"RoleArn"`
	Topic interface{} `yaml:"Topic"`
}

func (resource TopicRule_RepublishAction) Validate() []error {
	errs := []error{}
	
	
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.Topic == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Topic'"))
	}
	return errs
}
