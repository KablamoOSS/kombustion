package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// AppStreamImageBuilder Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html
type AppStreamImageBuilder struct {
	Type       string                          `yaml:"Type"`
	Properties AppStreamImageBuilderProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// AppStreamImageBuilder Properties
type AppStreamImageBuilderProperties struct {
	AppstreamAgentVersion       interface{} `yaml:"AppstreamAgentVersion,omitempty"`
	Description                 interface{} `yaml:"Description,omitempty"`
	DisplayName                 interface{} `yaml:"DisplayName,omitempty"`
	EnableDefaultInternetAccess interface{} `yaml:"EnableDefaultInternetAccess,omitempty"`
	ImageArn                    interface{} `yaml:"ImageArn,omitempty"`
	ImageName                   interface{} `yaml:"ImageName,omitempty"`
	InstanceType                interface{} `yaml:"InstanceType"`
	Name                        interface{} `yaml:"Name,omitempty"`
	VpcConfig                   interface{} `yaml:"VpcConfig,omitempty"`
	DomainJoinInfo              interface{} `yaml:"DomainJoinInfo,omitempty"`
}

// NewAppStreamImageBuilder constructor creates a new AppStreamImageBuilder
func NewAppStreamImageBuilder(properties AppStreamImageBuilderProperties, deps ...interface{}) AppStreamImageBuilder {
	return AppStreamImageBuilder{
		Type:       "AWS::AppStream::ImageBuilder",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppStreamImageBuilder parses AppStreamImageBuilder
func ParseAppStreamImageBuilder(
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
	var resource AppStreamImageBuilder
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
					"Fn::Sub": "${AWS::StackName}-AppStreamImageBuilder-" + name,
				},
			},
		},
	}

	return
}

// ParseAppStreamImageBuilder validator
func (resource AppStreamImageBuilder) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAppStreamImageBuilderProperties validator
func (resource AppStreamImageBuilderProperties) Validate() []error {
	errors := []error{}
	return errors
}