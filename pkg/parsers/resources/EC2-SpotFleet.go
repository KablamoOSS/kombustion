package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2SpotFleet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
type EC2SpotFleet struct {
	Type       string                 `yaml:"Type"`
	Properties EC2SpotFleetProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

// EC2SpotFleet Properties
type EC2SpotFleetProperties struct {
	SpotFleetRequestConfigData *properties.SpotFleetSpotFleetRequestConfigData `yaml:"SpotFleetRequestConfigData"`
}

// NewEC2SpotFleet constructor creates a new EC2SpotFleet
func NewEC2SpotFleet(properties EC2SpotFleetProperties, deps ...interface{}) EC2SpotFleet {
	return EC2SpotFleet{
		Type:       "AWS::EC2::SpotFleet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2SpotFleet parses EC2SpotFleet
func ParseEC2SpotFleet(
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
	var resource EC2SpotFleet
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

// ParseEC2SpotFleet validator
func (resource EC2SpotFleet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2SpotFleetProperties validator
func (resource EC2SpotFleetProperties) Validate() []error {
	errors := []error{}
	if resource.SpotFleetRequestConfigData == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SpotFleetRequestConfigData'"))
	} else {
		errors = append(errors, resource.SpotFleetRequestConfigData.Validate()...)
	}
	return errors
}
