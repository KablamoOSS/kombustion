package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationInput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html
type ApplicationInput struct {
	NamePrefix                   interface{}                              `yaml:"NamePrefix"`
	KinesisStreamsInput          *ApplicationKinesisStreamsInput          `yaml:"KinesisStreamsInput,omitempty"`
	KinesisFirehoseInput         *ApplicationKinesisFirehoseInput         `yaml:"KinesisFirehoseInput,omitempty"`
	InputSchema                  *ApplicationInputSchema                  `yaml:"InputSchema"`
	InputProcessingConfiguration *ApplicationInputProcessingConfiguration `yaml:"InputProcessingConfiguration,omitempty"`
	InputParallelism             *ApplicationInputParallelism             `yaml:"InputParallelism,omitempty"`
}

// ApplicationInput validation
func (resource ApplicationInput) Validate() []error {
	errors := []error{}

	if resource.NamePrefix == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'NamePrefix'"))
	}
	if resource.InputSchema == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'InputSchema'"))
	} else {
		errors = append(errors, resource.InputSchema.Validate()...)
	}
	return errors
}
