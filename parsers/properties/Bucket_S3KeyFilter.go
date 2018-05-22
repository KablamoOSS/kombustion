package properties

	import "fmt"

type Bucket_S3KeyFilter struct {
	
	Rules interface{} `yaml:"Rules"`
}

func (resource Bucket_S3KeyFilter) Validate() []error {
	errs := []error{}
	
	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
