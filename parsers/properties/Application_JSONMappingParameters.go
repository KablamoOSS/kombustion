package properties

	import "fmt"

type Application_JSONMappingParameters struct {
	
	RecordRowPath interface{} `yaml:"RecordRowPath"`
}

func (resource Application_JSONMappingParameters) Validate() []error {
	errs := []error{}
	
	if resource.RecordRowPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordRowPath'"))
	}
	return errs
}
