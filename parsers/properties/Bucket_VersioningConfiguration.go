package properties

	import "fmt"

type Bucket_VersioningConfiguration struct {
	
	Status interface{} `yaml:"Status"`
}

func (resource Bucket_VersioningConfiguration) Validate() []error {
	errs := []error{}
	
	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	return errs
}
