package properties

	import "fmt"

type ApplicationReferenceDataSource_S3ReferenceDataSource struct {
	
	
	
	BucketARN interface{} `yaml:"BucketARN"`
	FileKey interface{} `yaml:"FileKey"`
	ReferenceRoleARN interface{} `yaml:"ReferenceRoleARN"`
}

func (resource ApplicationReferenceDataSource_S3ReferenceDataSource) Validate() []error {
	errs := []error{}
	
	
	
	if resource.BucketARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketARN'"))
	}
	if resource.FileKey == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FileKey'"))
	}
	if resource.ReferenceRoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReferenceRoleARN'"))
	}
	return errs
}
