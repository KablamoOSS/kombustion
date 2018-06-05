package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Stack_ElasticIp struct {
	Ip   interface{} `yaml:"Ip"`
	Name interface{} `yaml:"Name,omitempty"`
}

func (resource Stack_ElasticIp) Validate() []error {
	errs := []error{}

	if resource.Ip == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Ip'"))
	}
	return errs
}
