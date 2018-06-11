package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TableAttributeDefinition Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
type TableAttributeDefinition struct {
	AttributeName interface{} `yaml:"AttributeName"`
	AttributeType interface{} `yaml:"AttributeType"`
}

func (resource TableAttributeDefinition) Validate() []error {
	errs := []error{}

	if resource.AttributeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeName'"))
	}
	if resource.AttributeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeType'"))
	}
	return errs
}
