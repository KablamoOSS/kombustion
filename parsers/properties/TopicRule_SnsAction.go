package properties

	import "fmt"

type TopicRule_SnsAction struct {
	
	
	
	MessageFormat interface{} `yaml:"MessageFormat,omitempty"`
	RoleArn interface{} `yaml:"RoleArn"`
	TargetArn interface{} `yaml:"TargetArn"`
}

func (resource TopicRule_SnsAction) Validate() []error {
	errs := []error{}
	
	
	
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.TargetArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetArn'"))
	}
	return errs
}
