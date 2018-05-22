package properties

	import "fmt"

type Table_StreamSpecification struct {
	
	StreamViewType interface{} `yaml:"StreamViewType"`
}

func (resource Table_StreamSpecification) Validate() []error {
	errs := []error{}
	
	if resource.StreamViewType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StreamViewType'"))
	}
	return errs
}
