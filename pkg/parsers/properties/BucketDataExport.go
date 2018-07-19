package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketDataExport Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html
type BucketDataExport struct {
	OutputSchemaVersion interface{}        `yaml:"OutputSchemaVersion"`
	Destination         *BucketDestination `yaml:"Destination"`
}

// BucketDataExport validation
func (resource BucketDataExport) Validate() []error {
	errors := []error{}

	if resource.OutputSchemaVersion == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'OutputSchemaVersion'"))
	}
	if resource.Destination == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Destination'"))
	} else {
		errors = append(errors, resource.Destination.Validate()...)
	}
	return errors
}
