package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Group_Policy struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	PolicyName     interface{} `yaml:"PolicyName"`
}

func (resource Group_Policy) Validate() []error {
	errs := []error{}

	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	return errs
}
