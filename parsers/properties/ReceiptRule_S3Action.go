package properties

	import "fmt"

type ReceiptRule_S3Action struct {
	
	
	
	
	BucketName interface{} `yaml:"BucketName"`
	KmsKeyArn interface{} `yaml:"KmsKeyArn,omitempty"`
	ObjectKeyPrefix interface{} `yaml:"ObjectKeyPrefix,omitempty"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

func (resource ReceiptRule_S3Action) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketName'"))
	}
	return errs
}
