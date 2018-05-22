package properties

	import "fmt"

type Bucket_BucketEncryption struct {
	
	ServerSideEncryptionConfiguration interface{} `yaml:"ServerSideEncryptionConfiguration"`
}

func (resource Bucket_BucketEncryption) Validate() []error {
	errs := []error{}
	
	if resource.ServerSideEncryptionConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServerSideEncryptionConfiguration'"))
	}
	return errs
}
