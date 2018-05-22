package properties

	import "fmt"

type ApplicationReferenceDataSource_CSVMappingParameters struct {
	
	
	RecordColumnDelimiter interface{} `yaml:"RecordColumnDelimiter"`
	RecordRowDelimiter interface{} `yaml:"RecordRowDelimiter"`
}

func (resource ApplicationReferenceDataSource_CSVMappingParameters) Validate() []error {
	errs := []error{}
	
	
	if resource.RecordColumnDelimiter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordColumnDelimiter'"))
	}
	if resource.RecordRowDelimiter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordRowDelimiter'"))
	}
	return errs
}
