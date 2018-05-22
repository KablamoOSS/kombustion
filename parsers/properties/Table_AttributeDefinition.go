package properties

	import "fmt"

type Table_AttributeDefinition struct {
	
	
	AttributeName interface{} `yaml:"AttributeName"`
	AttributeType interface{} `yaml:"AttributeType"`
}

func (resource Table_AttributeDefinition) Validate() []error {
	errs := []error{}
	
	
	if resource.AttributeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeName'"))
	}
	if resource.AttributeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeType'"))
	}
	return errs
}
