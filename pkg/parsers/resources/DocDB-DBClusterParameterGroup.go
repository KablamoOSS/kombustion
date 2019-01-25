package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DocDBDBClusterParameterGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html
type DocDBDBClusterParameterGroup struct {
	Type       string                                 `yaml:"Type"`
	Properties DocDBDBClusterParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// DocDBDBClusterParameterGroup Properties
type DocDBDBClusterParameterGroupProperties struct {
	Description interface{} `yaml:"Description"`
	Family      interface{} `yaml:"Family"`
	Name        interface{} `yaml:"Name,omitempty"`
	Parameters  interface{} `yaml:"Parameters"`
	Tags        interface{} `yaml:"Tags,omitempty"`
}

// NewDocDBDBClusterParameterGroup constructor creates a new DocDBDBClusterParameterGroup
func NewDocDBDBClusterParameterGroup(properties DocDBDBClusterParameterGroupProperties, deps ...interface{}) DocDBDBClusterParameterGroup {
	return DocDBDBClusterParameterGroup{
		Type:       "AWS::DocDB::DBClusterParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDocDBDBClusterParameterGroup parses DocDBDBClusterParameterGroup
func ParseDocDBDBClusterParameterGroup(
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
	var resource DocDBDBClusterParameterGroup
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
					"Fn::Sub": "${AWS::StackName}-DocDBDBClusterParameterGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseDocDBDBClusterParameterGroup validator
func (resource DocDBDBClusterParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDocDBDBClusterParameterGroupProperties validator
func (resource DocDBDBClusterParameterGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
