package properties

	import "fmt"

type TopicRule_PutItemInput struct {
	
	TableName interface{} `yaml:"TableName"`
}

func (resource TopicRule_PutItemInput) Validate() []error {
	errs := []error{}
	
	if resource.TableName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableName'"))
	}
	return errs
}
