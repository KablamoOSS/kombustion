package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// AutoScalingAutoScalingGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
type AutoScalingAutoScalingGroup struct {
	Type       string                                `yaml:"Type"`
	Properties AutoScalingAutoScalingGroupProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// AutoScalingAutoScalingGroup Properties
type AutoScalingAutoScalingGroupProperties struct {
	AutoScalingGroupName           interface{}                                             `yaml:"AutoScalingGroupName,omitempty"`
	Cooldown                       interface{}                                             `yaml:"Cooldown,omitempty"`
	DesiredCapacity                interface{}                                             `yaml:"DesiredCapacity,omitempty"`
	HealthCheckGracePeriod         interface{}                                             `yaml:"HealthCheckGracePeriod,omitempty"`
	HealthCheckType                interface{}                                             `yaml:"HealthCheckType,omitempty"`
	InstanceId                     interface{}                                             `yaml:"InstanceId,omitempty"`
	LaunchConfigurationName        interface{}                                             `yaml:"LaunchConfigurationName,omitempty"`
	MaxSize                        interface{}                                             `yaml:"MaxSize"`
	MinSize                        interface{}                                             `yaml:"MinSize"`
	PlacementGroup                 interface{}                                             `yaml:"PlacementGroup,omitempty"`
	ServiceLinkedRoleARN           interface{}                                             `yaml:"ServiceLinkedRoleARN,omitempty"`
	AvailabilityZones              interface{}                                             `yaml:"AvailabilityZones,omitempty"`
	TargetGroupARNs                interface{}                                             `yaml:"TargetGroupARNs,omitempty"`
	LifecycleHookSpecificationList interface{}                                             `yaml:"LifecycleHookSpecificationList,omitempty"`
	LoadBalancerNames              interface{}                                             `yaml:"LoadBalancerNames,omitempty"`
	MetricsCollection              interface{}                                             `yaml:"MetricsCollection,omitempty"`
	NotificationConfigurations     interface{}                                             `yaml:"NotificationConfigurations,omitempty"`
	Tags                           interface{}                                             `yaml:"Tags,omitempty"`
	TerminationPolicies            interface{}                                             `yaml:"TerminationPolicies,omitempty"`
	VPCZoneIdentifier              interface{}                                             `yaml:"VPCZoneIdentifier,omitempty"`
	LaunchTemplate                 *properties.AutoScalingGroupLaunchTemplateSpecification `yaml:"LaunchTemplate,omitempty"`
}

// NewAutoScalingAutoScalingGroup constructor creates a new AutoScalingAutoScalingGroup
func NewAutoScalingAutoScalingGroup(properties AutoScalingAutoScalingGroupProperties, deps ...interface{}) AutoScalingAutoScalingGroup {
	return AutoScalingAutoScalingGroup{
		Type:       "AWS::AutoScaling::AutoScalingGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAutoScalingAutoScalingGroup parses AutoScalingAutoScalingGroup
func ParseAutoScalingAutoScalingGroup(
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
	var resource AutoScalingAutoScalingGroup
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
					"Fn::Sub": "${AWS::StackName}-AutoScalingAutoScalingGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseAutoScalingAutoScalingGroup validator
func (resource AutoScalingAutoScalingGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAutoScalingAutoScalingGroupProperties validator
func (resource AutoScalingAutoScalingGroupProperties) Validate() []error {
	errors := []error{}
	if resource.MaxSize == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MaxSize'"))
	}
	if resource.MinSize == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MinSize'"))
	}
	return errors
}
