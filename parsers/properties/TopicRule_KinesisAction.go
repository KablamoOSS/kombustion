package properties

	import "fmt"

type TopicRule_KinesisAction struct {
	
	
	
	PartitionKey interface{} `yaml:"PartitionKey,omitempty"`
	RoleArn interface{} `yaml:"RoleArn"`
	StreamName interface{} `yaml:"StreamName"`
}

func (resource TopicRule_KinesisAction) Validate() []error {
	errs := []error{}
	
	
	
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.StreamName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StreamName'"))
	}
	return errs
}
