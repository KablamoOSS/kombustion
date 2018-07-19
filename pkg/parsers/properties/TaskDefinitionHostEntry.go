package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TaskDefinitionHostEntry Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html
type TaskDefinitionHostEntry struct {
	Hostname  interface{} `yaml:"Hostname"`
	IpAddress interface{} `yaml:"IpAddress"`
}

// TaskDefinitionHostEntry validation
func (resource TaskDefinitionHostEntry) Validate() []error {
	errors := []error{}

	if resource.Hostname == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Hostname'"))
	}
	if resource.IpAddress == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'IpAddress'"))
	}
	return errors
}
