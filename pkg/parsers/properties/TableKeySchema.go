package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TableKeySchema Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
type TableKeySchema struct {
	AttributeName interface{} `yaml:"AttributeName"`
	KeyType       interface{} `yaml:"KeyType"`
}

func (resource TableKeySchema) Validate() []error {
	errs := []error{}

	if resource.AttributeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeName'"))
	}
	if resource.KeyType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyType'"))
	}
	return errs
}
