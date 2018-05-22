package properties

	import "fmt"

type ApplicationReferenceDataSource_ReferenceSchema struct {
	
	
	
	RecordEncoding interface{} `yaml:"RecordEncoding,omitempty"`
	RecordFormat *ApplicationReferenceDataSource_RecordFormat `yaml:"RecordFormat"`
	RecordColumns interface{} `yaml:"RecordColumns"`
}

func (resource ApplicationReferenceDataSource_ReferenceSchema) Validate() []error {
	errs := []error{}
	
	
	
	if resource.RecordFormat == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordFormat'"))
	} else {
		errs = append(errs, resource.RecordFormat.Validate()...)
	}
	if resource.RecordColumns == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordColumns'"))
	}
	return errs
}
