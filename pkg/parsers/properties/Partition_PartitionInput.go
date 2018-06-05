package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Partition_PartitionInput struct {
	Parameters        interface{}                  `yaml:"Parameters,omitempty"`
	StorageDescriptor *Partition_StorageDescriptor `yaml:"StorageDescriptor,omitempty"`
	Values            interface{}                  `yaml:"Values"`
}

func (resource Partition_PartitionInput) Validate() []error {
	errs := []error{}

	if resource.Values == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Values'"))
	}
	return errs
}