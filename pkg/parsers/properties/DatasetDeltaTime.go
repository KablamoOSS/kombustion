package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DatasetDeltaTime Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatime.html
type DatasetDeltaTime struct {
	OffsetSeconds  interface{} `yaml:"OffsetSeconds"`
	TimeExpression interface{} `yaml:"TimeExpression"`
}

// DatasetDeltaTime validation
func (resource DatasetDeltaTime) Validate() []error {
	errors := []error{}

	return errors
}