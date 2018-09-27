package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SESConfigurationSetEventDestination Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html
type SESConfigurationSetEventDestination struct {
	Type       string                                        `yaml:"Type"`
	Properties SESConfigurationSetEventDestinationProperties `yaml:"Properties"`
	Condition  interface{}                                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                   `yaml:"DependsOn,omitempty"`
}

// SESConfigurationSetEventDestination Properties
type SESConfigurationSetEventDestinationProperties struct {
	ConfigurationSetName interface{}                                                  `yaml:"ConfigurationSetName"`
	EventDestination     *properties.ConfigurationSetEventDestinationEventDestination `yaml:"EventDestination"`
}

// NewSESConfigurationSetEventDestination constructor creates a new SESConfigurationSetEventDestination
func NewSESConfigurationSetEventDestination(properties SESConfigurationSetEventDestinationProperties, deps ...interface{}) SESConfigurationSetEventDestination {
	return SESConfigurationSetEventDestination{
		Type:       "AWS::SES::ConfigurationSetEventDestination",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSESConfigurationSetEventDestination parses SESConfigurationSetEventDestination
func ParseSESConfigurationSetEventDestination(
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

	// Resources
	var resource SESConfigurationSetEventDestination
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

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-SESConfigurationSetEventDestination-" + name,
				},
			},
		},
	}

	return
}

// ParseSESConfigurationSetEventDestination validator
func (resource SESConfigurationSetEventDestination) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSESConfigurationSetEventDestinationProperties validator
func (resource SESConfigurationSetEventDestinationProperties) Validate() []error {
	errors := []error{}
	if resource.ConfigurationSetName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ConfigurationSetName'"))
	}
	if resource.EventDestination == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'EventDestination'"))
	} else {
		errors = append(errors, resource.EventDestination.Validate()...)
	}
	return errors
}
