package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineParameterValue Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html
type PipelineParameterValue struct {
	Id          interface{} `yaml:"Id"`
	StringValue interface{} `yaml:"StringValue"`
}

func (resource PipelineParameterValue) Validate() []error {
	errs := []error{}

	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.StringValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StringValue'"))
	}
	return errs
}
