package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// IoTTopicRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
type IoTTopicRule struct {
	Type       string                 `yaml:"Type"`
	Properties IoTTopicRuleProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

// IoTTopicRule Properties
type IoTTopicRuleProperties struct {
	RuleName         interface{}                           `yaml:"RuleName,omitempty"`
	TopicRulePayload *properties.TopicRuleTopicRulePayload `yaml:"TopicRulePayload"`
}

// NewIoTTopicRule constructor creates a new IoTTopicRule
func NewIoTTopicRule(properties IoTTopicRuleProperties, deps ...interface{}) IoTTopicRule {
	return IoTTopicRule{
		Type:       "AWS::IoT::TopicRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIoTTopicRule parses IoTTopicRule
func ParseIoTTopicRule(
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
	var resource IoTTopicRule
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
					"Fn::Sub": "${AWS::StackName}-IoTTopicRule-" + name,
				},
			},
		},
	}

	return
}

// ParseIoTTopicRule validator
func (resource IoTTopicRule) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIoTTopicRuleProperties validator
func (resource IoTTopicRuleProperties) Validate() []error {
	errors := []error{}
	if resource.TopicRulePayload == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TopicRulePayload'"))
	} else {
		errors = append(errors, resource.TopicRulePayload.Validate()...)
	}
	return errors
}
