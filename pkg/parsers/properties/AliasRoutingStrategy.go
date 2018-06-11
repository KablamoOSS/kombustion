package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// AliasRoutingStrategy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html
type AliasRoutingStrategy struct {
	FleetId interface{} `yaml:"FleetId,omitempty"`
	Message interface{} `yaml:"Message,omitempty"`
	Type    interface{} `yaml:"Type"`
}

func (resource AliasRoutingStrategy) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
