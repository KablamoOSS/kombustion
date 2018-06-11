package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ComputeEnvironmentComputeResources Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html
type ComputeEnvironmentComputeResources struct {
	BidPercentage    interface{} `yaml:"BidPercentage,omitempty"`
	DesiredvCpus     interface{} `yaml:"DesiredvCpus,omitempty"`
	Ec2KeyPair       interface{} `yaml:"Ec2KeyPair,omitempty"`
	ImageId          interface{} `yaml:"ImageId,omitempty"`
	InstanceRole     interface{} `yaml:"InstanceRole"`
	MaxvCpus         interface{} `yaml:"MaxvCpus"`
	MinvCpus         interface{} `yaml:"MinvCpus"`
	SpotIamFleetRole interface{} `yaml:"SpotIamFleetRole,omitempty"`
	Tags             interface{} `yaml:"Tags,omitempty"`
	Type             interface{} `yaml:"Type"`
	InstanceTypes    interface{} `yaml:"InstanceTypes"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds"`
	Subnets          interface{} `yaml:"Subnets"`
}

func (resource ComputeEnvironmentComputeResources) Validate() []error {
	errs := []error{}

	if resource.InstanceRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceRole'"))
	}
	if resource.MaxvCpus == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxvCpus'"))
	}
	if resource.MinvCpus == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinvCpus'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.InstanceTypes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceTypes'"))
	}
	if resource.SecurityGroupIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityGroupIds'"))
	}
	if resource.Subnets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Subnets'"))
	}
	return errs
}
