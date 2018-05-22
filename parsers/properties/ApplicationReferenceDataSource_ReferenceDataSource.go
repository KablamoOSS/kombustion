package properties

	import "fmt"

type ApplicationReferenceDataSource_ReferenceDataSource struct {
	
	
	
	TableName interface{} `yaml:"TableName,omitempty"`
	S3ReferenceDataSource *ApplicationReferenceDataSource_S3ReferenceDataSource `yaml:"S3ReferenceDataSource,omitempty"`
	ReferenceSchema *ApplicationReferenceDataSource_ReferenceSchema `yaml:"ReferenceSchema"`
}

func (resource ApplicationReferenceDataSource_ReferenceDataSource) Validate() []error {
	errs := []error{}
	
	
	
	if resource.ReferenceSchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReferenceSchema'"))
	} else {
		errs = append(errs, resource.ReferenceSchema.Validate()...)
	}
	return errs
}
