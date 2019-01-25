package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SSMMaintenanceWindow Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html
type SSMMaintenanceWindow struct {
	Type       string                         `yaml:"Type"`
	Properties SSMMaintenanceWindowProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// SSMMaintenanceWindow Properties
type SSMMaintenanceWindowProperties struct {
	AllowUnassociatedTargets interface{} `yaml:"AllowUnassociatedTargets"`
	Cutoff                   interface{} `yaml:"Cutoff"`
	Description              interface{} `yaml:"Description,omitempty"`
	Duration                 interface{} `yaml:"Duration"`
	EndDate                  interface{} `yaml:"EndDate,omitempty"`
	Name                     interface{} `yaml:"Name"`
	Schedule                 interface{} `yaml:"Schedule"`
	ScheduleTimezone         interface{} `yaml:"ScheduleTimezone,omitempty"`
	StartDate                interface{} `yaml:"StartDate,omitempty"`
}

// NewSSMMaintenanceWindow constructor creates a new SSMMaintenanceWindow
func NewSSMMaintenanceWindow(properties SSMMaintenanceWindowProperties, deps ...interface{}) SSMMaintenanceWindow {
	return SSMMaintenanceWindow{
		Type:       "AWS::SSM::MaintenanceWindow",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSSMMaintenanceWindow parses SSMMaintenanceWindow
func ParseSSMMaintenanceWindow(
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
	var resource SSMMaintenanceWindow
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
					"Fn::Sub": "${AWS::StackName}-SSMMaintenanceWindow-" + name,
				},
			},
		},
	}

	return
}

// ParseSSMMaintenanceWindow validator
func (resource SSMMaintenanceWindow) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSSMMaintenanceWindowProperties validator
func (resource SSMMaintenanceWindowProperties) Validate() []error {
	errors := []error{}
	return errors
}