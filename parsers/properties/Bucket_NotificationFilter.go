package properties

	import "fmt"

type Bucket_NotificationFilter struct {
	
	S3Key *Bucket_S3KeyFilter `yaml:"S3Key"`
}

func (resource Bucket_NotificationFilter) Validate() []error {
	errs := []error{}
	
	if resource.S3Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Key'"))
	} else {
		errs = append(errs, resource.S3Key.Validate()...)
	}
	return errs
}
