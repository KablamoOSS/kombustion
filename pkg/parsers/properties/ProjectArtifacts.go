package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ProjectArtifacts Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html
type ProjectArtifacts struct {
	Location      interface{} `yaml:"Location,omitempty"`
	Name          interface{} `yaml:"Name,omitempty"`
	NamespaceType interface{} `yaml:"NamespaceType,omitempty"`
	Packaging     interface{} `yaml:"Packaging,omitempty"`
	Path          interface{} `yaml:"Path,omitempty"`
	Type          interface{} `yaml:"Type"`
}

func (resource ProjectArtifacts) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
