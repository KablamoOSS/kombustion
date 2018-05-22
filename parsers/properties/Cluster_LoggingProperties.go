package properties

	import "fmt"

type Cluster_LoggingProperties struct {
	
	
	BucketName interface{} `yaml:"BucketName"`
	S3KeyPrefix interface{} `yaml:"S3KeyPrefix,omitempty"`
}

func (resource Cluster_LoggingProperties) Validate() []error {
	errs := []error{}
	
	
	if resource.BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketName'"))
	}
	return errs
}
