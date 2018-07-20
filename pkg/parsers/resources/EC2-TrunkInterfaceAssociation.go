package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2TrunkInterfaceAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html
type EC2TrunkInterfaceAssociation struct {
	Type       string                                 `yaml:"Type"`
	Properties EC2TrunkInterfaceAssociationProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// EC2TrunkInterfaceAssociation Properties
type EC2TrunkInterfaceAssociationProperties struct {
	BranchInterfaceId interface{} `yaml:"BranchInterfaceId"`
	GREKey            interface{} `yaml:"GREKey,omitempty"`
	TrunkInterfaceId  interface{} `yaml:"TrunkInterfaceId"`
	VLANId            interface{} `yaml:"VLANId,omitempty"`
}

// NewEC2TrunkInterfaceAssociation constructor creates a new EC2TrunkInterfaceAssociation
func NewEC2TrunkInterfaceAssociation(properties EC2TrunkInterfaceAssociationProperties, deps ...interface{}) EC2TrunkInterfaceAssociation {
	return EC2TrunkInterfaceAssociation{
		Type:       "AWS::EC2::TrunkInterfaceAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2TrunkInterfaceAssociation parses EC2TrunkInterfaceAssociation
func ParseEC2TrunkInterfaceAssociation(
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
	var resource EC2TrunkInterfaceAssociation
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

// ParseEC2TrunkInterfaceAssociation validator
func (resource EC2TrunkInterfaceAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2TrunkInterfaceAssociationProperties validator
func (resource EC2TrunkInterfaceAssociationProperties) Validate() []error {
	errors := []error{}
	if resource.BranchInterfaceId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'BranchInterfaceId'"))
	}
	if resource.TrunkInterfaceId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TrunkInterfaceId'"))
	}
	return errors
}
