package properties

	import "fmt"

type Table_Order struct {
	
	
	Column interface{} `yaml:"Column"`
	SortOrder interface{} `yaml:"SortOrder"`
}

func (resource Table_Order) Validate() []error {
	errs := []error{}
	
	
	if resource.Column == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Column'"))
	}
	if resource.SortOrder == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SortOrder'"))
	}
	return errs
}
