package properties

	import "fmt"

type Bucket_LifecycleConfiguration struct {
	
	Rules interface{} `yaml:"Rules"`
}

func (resource Bucket_LifecycleConfiguration) Validate() []error {
	errs := []error{}
	
	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
