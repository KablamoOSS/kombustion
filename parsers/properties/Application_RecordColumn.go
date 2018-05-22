package properties

	import "fmt"

type Application_RecordColumn struct {
	
	
	
	Mapping interface{} `yaml:"Mapping,omitempty"`
	Name interface{} `yaml:"Name"`
	SqlType interface{} `yaml:"SqlType"`
}

func (resource Application_RecordColumn) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.SqlType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SqlType'"))
	}
	return errs
}
