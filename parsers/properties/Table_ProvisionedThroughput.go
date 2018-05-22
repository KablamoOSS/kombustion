package properties

	import "fmt"

type Table_ProvisionedThroughput struct {
	
	
	ReadCapacityUnits interface{} `yaml:"ReadCapacityUnits"`
	WriteCapacityUnits interface{} `yaml:"WriteCapacityUnits"`
}

func (resource Table_ProvisionedThroughput) Validate() []error {
	errs := []error{}
	
	
	if resource.ReadCapacityUnits == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReadCapacityUnits'"))
	}
	if resource.WriteCapacityUnits == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'WriteCapacityUnits'"))
	}
	return errs
}
