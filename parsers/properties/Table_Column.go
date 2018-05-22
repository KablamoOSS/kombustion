package properties

	import "fmt"

type Table_Column struct {
	
	
	
	Comment interface{} `yaml:"Comment,omitempty"`
	Name interface{} `yaml:"Name"`
	Type interface{} `yaml:"Type,omitempty"`
}

func (resource Table_Column) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
