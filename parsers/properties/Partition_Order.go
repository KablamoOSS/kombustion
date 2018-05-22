package properties

	import "fmt"

type Partition_Order struct {
	
	
	Column interface{} `yaml:"Column"`
	SortOrder interface{} `yaml:"SortOrder,omitempty"`
}

func (resource Partition_Order) Validate() []error {
	errs := []error{}
	
	
	if resource.Column == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Column'"))
	}
	return errs
}
