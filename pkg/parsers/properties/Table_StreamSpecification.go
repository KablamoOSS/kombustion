package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Table_StreamSpecification struct {
	StreamViewType interface{} `yaml:"StreamViewType"`
}

func (resource Table_StreamSpecification) Validate() []error {
	errs := []error{}

	if resource.StreamViewType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StreamViewType'"))
	}
	return errs
}