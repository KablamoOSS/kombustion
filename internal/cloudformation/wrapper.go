package cloudformation

import (
	"fmt"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Subset of AWS API calls required by the events task, as a mockable
// interface.
type StackEventer interface {
	Open(string, string) string
	DescribeStackEvents(*awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error)
	DescribeStackEventsPages(*awsCF.DescribeStackEventsInput, func(*awsCF.DescribeStackEventsOutput, bool) bool) error
}

// Subset of AWS API calls required by the upsert task, as a mockable
// interface.
type StackUpserter interface {
	Open(string, string) string
	CreateChangeSet(*awsCF.CreateChangeSetInput) (*awsCF.CreateChangeSetOutput, error)
	DeleteChangeSet(*awsCF.DeleteChangeSetInput) (*awsCF.DeleteChangeSetOutput, error)
	DeleteStack(*awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error)
	DescribeChangeSet(*awsCF.DescribeChangeSetInput) (*awsCF.DescribeChangeSetOutput, error)
	DescribeStackEvents(*awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error)
	DescribeStackEventsPages(*awsCF.DescribeStackEventsInput, func(*awsCF.DescribeStackEventsOutput, bool) bool) error
	DescribeStacks(*awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error)
	ExecuteChangeSet(*awsCF.ExecuteChangeSetInput) (*awsCF.ExecuteChangeSetOutput, error)
	WaitUntilChangeSetCreateComplete(*awsCF.DescribeChangeSetInput) error
}

type StackDeleter interface {
	Open(string, string) string
	DeleteStack(*awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error)
	DescribeStackEvents(*awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error)
	DescribeStackEventsPages(*awsCF.DescribeStackEventsInput, func(*awsCF.DescribeStackEventsOutput, bool) bool) error
	DescribeStacks(*awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error)
}

type Wrapper struct {
	client *awsCF.CloudFormation
}

// Placeholder helper so that tasks still using awsCF.CloudFormation directly
// can wrap it to use helpers that have been updated to use interfaces.
func NewWrapper(client *awsCF.CloudFormation) *Wrapper {
	return &Wrapper{
		client: client,
	}
}

func (wr *Wrapper) Open(profile, region string) string {
	acctID, cfClient := GetCloudformationClient(profile, region)
	wr.client = cfClient
	return acctID
}

func (wr *Wrapper) DescribeStackEvents(input *awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DescribeStackEvents(input)
}

func (wr *Wrapper) DescribeStackEventsPages(input *awsCF.DescribeStackEventsInput, fn func(*awsCF.DescribeStackEventsOutput, bool) bool) error {
	if wr.client == nil {
		return fmt.Errorf("connection not opened")
	}
	return wr.client.DescribeStackEventsPages(input, fn)
}

func (wr *Wrapper) DescribeStacks(in *awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DescribeStacks(in)
}

func (wr *Wrapper) DeleteStack(in *awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DeleteStack(in)
}

func (wr *Wrapper) CreateChangeSet(in *awsCF.CreateChangeSetInput) (*awsCF.CreateChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.CreateChangeSet(in)
}

func (wr *Wrapper) DescribeChangeSet(in *awsCF.DescribeChangeSetInput) (*awsCF.DescribeChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DescribeChangeSet(in)
}

func (wr *Wrapper) WaitUntilChangeSetCreateComplete(in *awsCF.DescribeChangeSetInput) error {
	if wr.client == nil {
		return fmt.Errorf("connection not opened")
	}
	return wr.client.WaitUntilChangeSetCreateComplete(in)
}

func (wr *Wrapper) ExecuteChangeSet(in *awsCF.ExecuteChangeSetInput) (*awsCF.ExecuteChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.ExecuteChangeSet(in)
}

func (wr *Wrapper) DeleteChangeSet(in *awsCF.DeleteChangeSetInput) (*awsCF.DeleteChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DeleteChangeSet(in)
}
