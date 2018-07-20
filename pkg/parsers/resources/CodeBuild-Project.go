package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CodeBuildProject Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
type CodeBuildProject struct {
	Type       string                     `yaml:"Type"`
	Properties CodeBuildProjectProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// CodeBuildProject Properties
type CodeBuildProjectProperties struct {
	BadgeEnabled     interface{}                        `yaml:"BadgeEnabled,omitempty"`
	Description      interface{}                        `yaml:"Description,omitempty"`
	EncryptionKey    interface{}                        `yaml:"EncryptionKey,omitempty"`
	Name             interface{}                        `yaml:"Name,omitempty"`
	ServiceRole      interface{}                        `yaml:"ServiceRole"`
	TimeoutInMinutes interface{}                        `yaml:"TimeoutInMinutes,omitempty"`
	VpcConfig        *properties.ProjectVpcConfig       `yaml:"VpcConfig,omitempty"`
	Source           *properties.ProjectSource          `yaml:"Source"`
	Triggers         *properties.ProjectProjectTriggers `yaml:"Triggers,omitempty"`
	Cache            *properties.ProjectProjectCache    `yaml:"Cache,omitempty"`
	Tags             interface{}                        `yaml:"Tags,omitempty"`
	Environment      *properties.ProjectEnvironment     `yaml:"Environment"`
	Artifacts        *properties.ProjectArtifacts       `yaml:"Artifacts"`
}

// NewCodeBuildProject constructor creates a new CodeBuildProject
func NewCodeBuildProject(properties CodeBuildProjectProperties, deps ...interface{}) CodeBuildProject {
	return CodeBuildProject{
		Type:       "AWS::CodeBuild::Project",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodeBuildProject parses CodeBuildProject
func ParseCodeBuildProject(
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
	var resource CodeBuildProject
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

// ParseCodeBuildProject validator
func (resource CodeBuildProject) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCodeBuildProjectProperties validator
func (resource CodeBuildProjectProperties) Validate() []error {
	errors := []error{}
	if resource.ServiceRole == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ServiceRole'"))
	}
	if resource.Source == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Source'"))
	} else {
		errors = append(errors, resource.Source.Validate()...)
	}
	if resource.Environment == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Environment'"))
	} else {
		errors = append(errors, resource.Environment.Validate()...)
	}
	if resource.Artifacts == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Artifacts'"))
	} else {
		errors = append(errors, resource.Artifacts.Validate()...)
	}
	return errors
}
