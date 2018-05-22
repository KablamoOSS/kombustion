package properties

	import "fmt"

type Table_SSESpecification struct {
	
	SSEEnabled interface{} `yaml:"SSEEnabled"`
}

func (resource Table_SSESpecification) Validate() []error {
	errs := []error{}
	
	if resource.SSEEnabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SSEEnabled'"))
	}
	return errs
}
