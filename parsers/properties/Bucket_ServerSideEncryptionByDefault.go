package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Bucket_ServerSideEncryptionByDefault struct {
	KMSMasterKeyID interface{} `yaml:"KMSMasterKeyID,omitempty"`
	SSEAlgorithm   interface{} `yaml:"SSEAlgorithm"`
}

func (resource Bucket_ServerSideEncryptionByDefault) Validate() []error {
	errs := []error{}

	if resource.SSEAlgorithm == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SSEAlgorithm'"))
	}
	return errs
}
