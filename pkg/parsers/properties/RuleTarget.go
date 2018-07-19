package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// RuleTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html
type RuleTarget struct {
	Arn                  interface{}               `yaml:"Arn"`
	Id                   interface{}               `yaml:"Id"`
	Input                interface{}               `yaml:"Input,omitempty"`
	InputPath            interface{}               `yaml:"InputPath,omitempty"`
	RoleArn              interface{}               `yaml:"RoleArn,omitempty"`
	RunCommandParameters *RuleRunCommandParameters `yaml:"RunCommandParameters,omitempty"`
	KinesisParameters    *RuleKinesisParameters    `yaml:"KinesisParameters,omitempty"`
	InputTransformer     *RuleInputTransformer     `yaml:"InputTransformer,omitempty"`
	EcsParameters        *RuleEcsParameters        `yaml:"EcsParameters,omitempty"`
}

// RuleTarget validation
func (resource RuleTarget) Validate() []error {
	errors := []error{}

	if resource.Arn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Arn'"))
	}
	if resource.Id == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Id'"))
	}
	return errors
}
