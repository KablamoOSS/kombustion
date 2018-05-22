package properties

	import "fmt"

type LoadBalancer_AccessLoggingPolicy struct {
	
	
	
	
	EmitInterval interface{} `yaml:"EmitInterval,omitempty"`
	Enabled interface{} `yaml:"Enabled"`
	S3BucketName interface{} `yaml:"S3BucketName"`
	S3BucketPrefix interface{} `yaml:"S3BucketPrefix,omitempty"`
}

func (resource LoadBalancer_AccessLoggingPolicy) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	if resource.S3BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3BucketName'"))
	}
	return errs
}
