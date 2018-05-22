package properties

	import "fmt"

type TopicRule_S3Action struct {
	
	
	
	BucketName interface{} `yaml:"BucketName"`
	Key interface{} `yaml:"Key"`
	RoleArn interface{} `yaml:"RoleArn"`
}

func (resource TopicRule_S3Action) Validate() []error {
	errs := []error{}
	
	
	
	if resource.BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketName'"))
	}
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
