package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Bucket_TagFilter struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

func (resource Bucket_TagFilter) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
