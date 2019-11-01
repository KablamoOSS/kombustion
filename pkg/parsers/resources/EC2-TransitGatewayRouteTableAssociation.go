package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2TransitGatewayRouteTableAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html
type EC2TransitGatewayRouteTableAssociation struct {
	Type       string                                           `yaml:"Type"`
	Properties EC2TransitGatewayRouteTableAssociationProperties `yaml:"Properties"`
	Condition  interface{}                                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                      `yaml:"DependsOn,omitempty"`
}

// EC2TransitGatewayRouteTableAssociation Properties
type EC2TransitGatewayRouteTableAssociationProperties struct {
	TransitGatewayAttachmentId interface{} `yaml:"TransitGatewayAttachmentId"`
	TransitGatewayRouteTableId interface{} `yaml:"TransitGatewayRouteTableId"`
}

// NewEC2TransitGatewayRouteTableAssociation constructor creates a new EC2TransitGatewayRouteTableAssociation
func NewEC2TransitGatewayRouteTableAssociation(properties EC2TransitGatewayRouteTableAssociationProperties, deps ...interface{}) EC2TransitGatewayRouteTableAssociation {
	return EC2TransitGatewayRouteTableAssociation{
		Type:       "AWS::EC2::TransitGatewayRouteTableAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2TransitGatewayRouteTableAssociation parses EC2TransitGatewayRouteTableAssociation
func ParseEC2TransitGatewayRouteTableAssociation(
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
	var resource EC2TransitGatewayRouteTableAssociation
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
					"Fn::Sub": "${AWS::StackName}-EC2TransitGatewayRouteTableAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2TransitGatewayRouteTableAssociation validator
func (resource EC2TransitGatewayRouteTableAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2TransitGatewayRouteTableAssociationProperties validator
func (resource EC2TransitGatewayRouteTableAssociationProperties) Validate() []error {
	errors := []error{}
	return errors
}