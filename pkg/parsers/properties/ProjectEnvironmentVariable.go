package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ProjectEnvironmentVariable Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html
type ProjectEnvironmentVariable struct {
	Name  interface{} `yaml:"Name"`
	Type  interface{} `yaml:"Type,omitempty"`
	Value interface{} `yaml:"Value"`
}

func (resource ProjectEnvironmentVariable) Validate() []error {
	errs := []error{}

	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
