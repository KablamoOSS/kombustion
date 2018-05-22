package properties

	import "fmt"

type Table_GlobalSecondaryIndex struct {
	
	
	
	
	IndexName interface{} `yaml:"IndexName"`
	ProvisionedThroughput *Table_ProvisionedThroughput `yaml:"ProvisionedThroughput"`
	Projection *Table_Projection `yaml:"Projection"`
	KeySchema interface{} `yaml:"KeySchema"`
}

func (resource Table_GlobalSecondaryIndex) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.IndexName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IndexName'"))
	}
	if resource.ProvisionedThroughput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ProvisionedThroughput'"))
	} else {
		errs = append(errs, resource.ProvisionedThroughput.Validate()...)
	}
	if resource.Projection == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Projection'"))
	} else {
		errs = append(errs, resource.Projection.Validate()...)
	}
	if resource.KeySchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeySchema'"))
	}
	return errs
}
