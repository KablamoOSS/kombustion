package properties

	import "fmt"

type DeliveryStream_CopyCommand struct {
	
	
	
	CopyOptions interface{} `yaml:"CopyOptions,omitempty"`
	DataTableColumns interface{} `yaml:"DataTableColumns,omitempty"`
	DataTableName interface{} `yaml:"DataTableName"`
}

func (resource DeliveryStream_CopyCommand) Validate() []error {
	errs := []error{}
	
	
	
	if resource.DataTableName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DataTableName'"))
	}
	return errs
}
