package properties

	import "fmt"

type Bucket_DataExport struct {
	
	
	OutputSchemaVersion interface{} `yaml:"OutputSchemaVersion"`
	Destination *Bucket_Destination `yaml:"Destination"`
}

func (resource Bucket_DataExport) Validate() []error {
	errs := []error{}
	
	
	if resource.OutputSchemaVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OutputSchemaVersion'"))
	}
	if resource.Destination == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Destination'"))
	} else {
		errs = append(errs, resource.Destination.Validate()...)
	}
	return errs
}
