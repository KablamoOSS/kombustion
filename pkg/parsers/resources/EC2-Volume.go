package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2Volume Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html
type EC2Volume struct {
	Type       string              `yaml:"Type"`
	Properties EC2VolumeProperties `yaml:"Properties"`
	Condition  interface{}         `yaml:"Condition,omitempty"`
	Metadata   interface{}         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}         `yaml:"DependsOn,omitempty"`
}

// EC2Volume Properties
type EC2VolumeProperties struct {
	AutoEnableIO     interface{} `yaml:"AutoEnableIO,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone"`
	Encrypted        interface{} `yaml:"Encrypted,omitempty"`
	Iops             interface{} `yaml:"Iops,omitempty"`
	KmsKeyId         interface{} `yaml:"KmsKeyId,omitempty"`
	Size             interface{} `yaml:"Size,omitempty"`
	SnapshotId       interface{} `yaml:"SnapshotId,omitempty"`
	VolumeType       interface{} `yaml:"VolumeType,omitempty"`
	Tags             interface{} `yaml:"Tags,omitempty"`
}

// NewEC2Volume constructor creates a new EC2Volume
func NewEC2Volume(properties EC2VolumeProperties, deps ...interface{}) EC2Volume {
	return EC2Volume{
		Type:       "AWS::EC2::Volume",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2Volume parses EC2Volume
func ParseEC2Volume(
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
	var resource EC2Volume
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

// ParseEC2Volume validator
func (resource EC2Volume) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VolumeProperties validator
func (resource EC2VolumeProperties) Validate() []error {
	errors := []error{}
	if resource.AvailabilityZone == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'AvailabilityZone'"))
	}
	return errors
}
