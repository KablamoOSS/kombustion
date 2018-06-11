package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ServicePlacementConstraint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html
type ServicePlacementConstraint struct {
	Expression interface{} `yaml:"Expression,omitempty"`
	Type       interface{} `yaml:"Type"`
}

func (resource ServicePlacementConstraint) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
