package properties

	import "fmt"

type ApplicationVersion_SourceBundle struct {
	
	
	S3Bucket interface{} `yaml:"S3Bucket"`
	S3Key interface{} `yaml:"S3Key"`
}

func (resource ApplicationVersion_SourceBundle) Validate() []error {
	errs := []error{}
	
	
	if resource.S3Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Bucket'"))
	}
	if resource.S3Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Key'"))
	}
	return errs
}
