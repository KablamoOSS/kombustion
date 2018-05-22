package properties

	import "fmt"

type Table_KeySchema struct {
	
	
	AttributeName interface{} `yaml:"AttributeName"`
	KeyType interface{} `yaml:"KeyType"`
}

func (resource Table_KeySchema) Validate() []error {
	errs := []error{}
	
	
	if resource.AttributeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeName'"))
	}
	if resource.KeyType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyType'"))
	}
	return errs
}
