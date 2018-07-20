package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// GlueJob Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html
type GlueJob struct {
	Type       string            `yaml:"Type"`
	Properties GlueJobProperties `yaml:"Properties"`
	Condition  interface{}       `yaml:"Condition,omitempty"`
	Metadata   interface{}       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}       `yaml:"DependsOn,omitempty"`
}

// GlueJob Properties
type GlueJobProperties struct {
	AllocatedCapacity interface{}                      `yaml:"AllocatedCapacity,omitempty"`
	DefaultArguments  interface{}                      `yaml:"DefaultArguments,omitempty"`
	Description       interface{}                      `yaml:"Description,omitempty"`
	LogUri            interface{}                      `yaml:"LogUri,omitempty"`
	MaxRetries        interface{}                      `yaml:"MaxRetries,omitempty"`
	Name              interface{}                      `yaml:"Name,omitempty"`
	Role              interface{}                      `yaml:"Role"`
	Command           *properties.JobJobCommand        `yaml:"Command"`
	ExecutionProperty *properties.JobExecutionProperty `yaml:"ExecutionProperty,omitempty"`
	Connections       *properties.JobConnectionsList   `yaml:"Connections,omitempty"`
}

// NewGlueJob constructor creates a new GlueJob
func NewGlueJob(properties GlueJobProperties, deps ...interface{}) GlueJob {
	return GlueJob{
		Type:       "AWS::Glue::Job",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueJob parses GlueJob
func ParseGlueJob(
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
	var resource GlueJob
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

// ParseGlueJob validator
func (resource GlueJob) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueJobProperties validator
func (resource GlueJobProperties) Validate() []error {
	errors := []error{}
	if resource.Role == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Command == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Command'"))
	} else {
		errors = append(errors, resource.Command.Validate()...)
	}
	return errors
}
