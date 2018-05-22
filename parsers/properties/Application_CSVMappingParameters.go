package properties

	import "fmt"

type Application_CSVMappingParameters struct {
	
	
	RecordColumnDelimiter interface{} `yaml:"RecordColumnDelimiter"`
	RecordRowDelimiter interface{} `yaml:"RecordRowDelimiter"`
}

func (resource Application_CSVMappingParameters) Validate() []error {
	errs := []error{}
	
	
	if resource.RecordColumnDelimiter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordColumnDelimiter'"))
	}
	if resource.RecordRowDelimiter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordRowDelimiter'"))
	}
	return errs
}
