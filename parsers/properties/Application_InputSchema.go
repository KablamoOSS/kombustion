package properties

	import "fmt"

type Application_InputSchema struct {
	
	
	
	RecordEncoding interface{} `yaml:"RecordEncoding,omitempty"`
	RecordFormat *Application_RecordFormat `yaml:"RecordFormat"`
	RecordColumns interface{} `yaml:"RecordColumns"`
}

func (resource Application_InputSchema) Validate() []error {
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
