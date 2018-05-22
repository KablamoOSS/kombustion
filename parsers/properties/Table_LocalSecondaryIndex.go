package properties

	import "fmt"

type Table_LocalSecondaryIndex struct {
	
	
	
	IndexName interface{} `yaml:"IndexName"`
	Projection *Table_Projection `yaml:"Projection"`
	KeySchema interface{} `yaml:"KeySchema"`
}

func (resource Table_LocalSecondaryIndex) Validate() []error {
	errs := []error{}
	
	
	
	if resource.IndexName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IndexName'"))
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
