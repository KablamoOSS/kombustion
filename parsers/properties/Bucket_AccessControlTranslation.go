package properties

	import "fmt"

type Bucket_AccessControlTranslation struct {
	
	Owner interface{} `yaml:"Owner"`
}

func (resource Bucket_AccessControlTranslation) Validate() []error {
	errs := []error{}
	
	if resource.Owner == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Owner'"))
	}
	return errs
}
