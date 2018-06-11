package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TableProjection Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
type TableProjection struct {
	ProjectionType   interface{} `yaml:"ProjectionType,omitempty"`
	NonKeyAttributes interface{} `yaml:"NonKeyAttributes,omitempty"`
}

func (resource TableProjection) Validate() []error {
	errs := []error{}

	return errs
}
