package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// KinesisFirehoseDeliveryStream Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
type KinesisFirehoseDeliveryStream struct {
	Type       string                                  `yaml:"Type"`
	Properties KinesisFirehoseDeliveryStreamProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

// KinesisFirehoseDeliveryStream Properties
type KinesisFirehoseDeliveryStreamProperties struct {
	DeliveryStreamName                    interface{}                                                     `yaml:"DeliveryStreamName,omitempty"`
	DeliveryStreamType                    interface{}                                                     `yaml:"DeliveryStreamType,omitempty"`
	SplunkDestinationConfiguration        *properties.DeliveryStreamSplunkDestinationConfiguration        `yaml:"SplunkDestinationConfiguration,omitempty"`
	S3DestinationConfiguration            *properties.DeliveryStreamS3DestinationConfiguration            `yaml:"S3DestinationConfiguration,omitempty"`
	RedshiftDestinationConfiguration      *properties.DeliveryStreamRedshiftDestinationConfiguration      `yaml:"RedshiftDestinationConfiguration,omitempty"`
	KinesisStreamSourceConfiguration      *properties.DeliveryStreamKinesisStreamSourceConfiguration      `yaml:"KinesisStreamSourceConfiguration,omitempty"`
	ExtendedS3DestinationConfiguration    *properties.DeliveryStreamExtendedS3DestinationConfiguration    `yaml:"ExtendedS3DestinationConfiguration,omitempty"`
	ElasticsearchDestinationConfiguration *properties.DeliveryStreamElasticsearchDestinationConfiguration `yaml:"ElasticsearchDestinationConfiguration,omitempty"`
}

// NewKinesisFirehoseDeliveryStream constructor creates a new KinesisFirehoseDeliveryStream
func NewKinesisFirehoseDeliveryStream(properties KinesisFirehoseDeliveryStreamProperties, deps ...interface{}) KinesisFirehoseDeliveryStream {
	return KinesisFirehoseDeliveryStream{
		Type:       "AWS::KinesisFirehose::DeliveryStream",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseKinesisFirehoseDeliveryStream parses KinesisFirehoseDeliveryStream
func ParseKinesisFirehoseDeliveryStream(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource KinesisFirehoseDeliveryStream
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseKinesisFirehoseDeliveryStream validator
func (resource KinesisFirehoseDeliveryStream) Validate() []error {
	return resource.Properties.Validate()
}

// ParseKinesisFirehoseDeliveryStreamProperties validator
func (resource KinesisFirehoseDeliveryStreamProperties) Validate() []error {
	errors := []error{}
	return errors
}
