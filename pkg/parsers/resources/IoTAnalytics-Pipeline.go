package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// IoTAnalyticsPipeline Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-pipeline.html
type IoTAnalyticsPipeline struct {
	Type       string                         `yaml:"Type"`
	Properties IoTAnalyticsPipelineProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// IoTAnalyticsPipeline Properties
type IoTAnalyticsPipelineProperties struct {
	PipelineName       interface{} `yaml:"PipelineName,omitempty"`
	PipelineActivities interface{} `yaml:"PipelineActivities"`
	Tags               interface{} `yaml:"Tags,omitempty"`
}

// NewIoTAnalyticsPipeline constructor creates a new IoTAnalyticsPipeline
func NewIoTAnalyticsPipeline(properties IoTAnalyticsPipelineProperties, deps ...interface{}) IoTAnalyticsPipeline {
	return IoTAnalyticsPipeline{
		Type:       "AWS::IoTAnalytics::Pipeline",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIoTAnalyticsPipeline parses IoTAnalyticsPipeline
func ParseIoTAnalyticsPipeline(
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
	var resource IoTAnalyticsPipeline
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
					"Fn::Sub": "${AWS::StackName}-IoTAnalyticsPipeline-" + name,
				},
			},
		},
	}

	return
}

// ParseIoTAnalyticsPipeline validator
func (resource IoTAnalyticsPipeline) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIoTAnalyticsPipelineProperties validator
func (resource IoTAnalyticsPipelineProperties) Validate() []error {
	errors := []error{}
	return errors
}
