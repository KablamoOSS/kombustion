package properties

	import "fmt"

type Bucket_AbortIncompleteMultipartUpload struct {
	
	DaysAfterInitiation interface{} `yaml:"DaysAfterInitiation"`
}

func (resource Bucket_AbortIncompleteMultipartUpload) Validate() []error {
	errs := []error{}
	
	if resource.DaysAfterInitiation == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DaysAfterInitiation'"))
	}
	return errs
}
