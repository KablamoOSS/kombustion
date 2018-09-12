package cloudformation

import (
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

type StackEventer interface {
	Open(string, string) string
	DescribeStackEvents(*awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error)
	DescribeStackEventsPages(*awsCF.DescribeStackEventsInput, func(*awsCF.DescribeStackEventsOutput, bool) bool) error
}

type Wrapper struct {
	client *awsCF.CloudFormation
}

func NewWrapper(client *awsCF.CloudFormation) *Wrapper {
	return &Wrapper{
		client: client,
	}
}

func (cfe *Wrapper) Open(profile, region string) string {
	acctID, cfClient := GetCloudformationClient(profile, region)
	cfe.client = cfClient
	return acctID
}

func (cfe *Wrapper) DescribeStackEvents(input *awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error) {
	return cfe.client.DescribeStackEvents(input)
}

func (cfe *Wrapper) DescribeStackEventsPages(input *awsCF.DescribeStackEventsInput, fn func(*awsCF.DescribeStackEventsOutput, bool) bool) error {
	return cfe.client.DescribeStackEventsPages(input, fn)
}
