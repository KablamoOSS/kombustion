package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Pipeline_EncryptionKey struct {
	Id   interface{} `yaml:"Id"`
	Type interface{} `yaml:"Type"`
}

func (resource Pipeline_EncryptionKey) Validate() []error {
	errs := []error{}

	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
