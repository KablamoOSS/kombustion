package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type ConfigRule_Source struct {
	Owner            interface{} `yaml:"Owner"`
	SourceIdentifier interface{} `yaml:"SourceIdentifier"`
	SourceDetails    interface{} `yaml:"SourceDetails,omitempty"`
}

func (resource ConfigRule_Source) Validate() []error {
	errs := []error{}

	if resource.Owner == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Owner'"))
	}
	if resource.SourceIdentifier == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SourceIdentifier'"))
	}
	return errs
}
