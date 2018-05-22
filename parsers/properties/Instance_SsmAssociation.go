package properties

	import "fmt"

type Instance_SsmAssociation struct {
	
	
	DocumentName interface{} `yaml:"DocumentName"`
	AssociationParameters interface{} `yaml:"AssociationParameters,omitempty"`
}

func (resource Instance_SsmAssociation) Validate() []error {
	errs := []error{}
	
	
	if resource.DocumentName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DocumentName'"))
	}
	return errs
}
