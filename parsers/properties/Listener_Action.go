package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Listener_Action struct {
	TargetGroupArn interface{} `yaml:"TargetGroupArn"`
	Type           interface{} `yaml:"Type"`
}

func (resource Listener_Action) Validate() []error {
	errs := []error{}

	if resource.TargetGroupArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetGroupArn'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
