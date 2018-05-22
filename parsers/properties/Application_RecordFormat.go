package properties

	import "fmt"

type Application_RecordFormat struct {
	
	
	RecordFormatType interface{} `yaml:"RecordFormatType"`
	MappingParameters *Application_MappingParameters `yaml:"MappingParameters,omitempty"`
}

func (resource Application_RecordFormat) Validate() []error {
	errs := []error{}
	
	
	if resource.RecordFormatType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordFormatType'"))
	}
	return errs
}
