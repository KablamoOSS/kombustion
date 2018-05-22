package properties

	import "fmt"

type Bucket_NoncurrentVersionTransition struct {
	
	
	StorageClass interface{} `yaml:"StorageClass"`
	TransitionInDays interface{} `yaml:"TransitionInDays"`
}

func (resource Bucket_NoncurrentVersionTransition) Validate() []error {
	errs := []error{}
	
	
	if resource.StorageClass == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StorageClass'"))
	}
	if resource.TransitionInDays == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TransitionInDays'"))
	}
	return errs
}
