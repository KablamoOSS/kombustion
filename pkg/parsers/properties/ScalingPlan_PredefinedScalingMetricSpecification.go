package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type ScalingPlan_PredefinedScalingMetricSpecification struct {
	PredefinedScalingMetricType interface{} `yaml:"PredefinedScalingMetricType"`
	ResourceLabel               interface{} `yaml:"ResourceLabel,omitempty"`
}

func (resource ScalingPlan_PredefinedScalingMetricSpecification) Validate() []error {
	errs := []error{}

	if resource.PredefinedScalingMetricType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PredefinedScalingMetricType'"))
	}
	return errs
}