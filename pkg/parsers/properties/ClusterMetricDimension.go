package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterMetricDimension Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html
type ClusterMetricDimension struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

// ClusterMetricDimension validation
func (resource ClusterMetricDimension) Validate() []error {
	errors := []error{}

	if resource.Key == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Value'"))
	}
	return errors
}
