package tasks

import (
	"fmt"
	"testing"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/coretest"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

type MockStackDeleter struct {
	AcctID string
	Events  map[string][]*awsCF.StackEvent
	Stacks  map[string]*awsCF.Stack
}


func (msd *MockStackDeleter) Open(_, _ string) string {
	if msd.Events == nil {
		msd.Events = make(map[string][]*awsCF.StackEvent)
	}
	if msd.Stacks == nil {
		msd.Stacks = make(map[string]*awsCF.Stack)
	}
	return msd.AcctID
}

func (msd *MockStackDeleter) DeleteStack(input *awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error) {
	stack, ok := msd.Stacks[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("stack not found: %s", *input.StackName)
	}
	stack.StackStatus = aws.String("DELETE_IN_PROGRESS")
	return &awsCF.DeleteStackOutput{}, nil
}

func (msd *MockStackDeleter) DescribeStackEvents(input *awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error) {
	events, ok := msd.Events[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("stack %s not found", *input.StackName)
	}

	out := &awsCF.DescribeStackEventsOutput{
		StackEvents: events,
	}

	return out, nil
}

func (msd *MockStackDeleter) DescribeStackEventsPages(input *awsCF.DescribeStackEventsInput, fn func(*awsCF.DescribeStackEventsOutput, bool) bool) error {
	out, err := msd.DescribeStackEvents(input)
	if err != nil {
		return err
	}

	fn(out, true)
	return nil
}

func (msd *MockStackDeleter) DescribeStacks(input *awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error) {
	stack, ok := msd.Stacks[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("stack not found: %s", *input.StackName)
	}
	output := &awsCF.DescribeStacksOutput{
		Stacks: []*awsCF.Stack{stack},
	}
	if *stack.StackStatus == "DELETE_IN_PROGRESS" {
		delete(msd.Stacks, *input.StackName)
		delete(msd.Events, *input.StackName)
	}
	return output, nil
}

func TestDeleteTask(t *testing.T) {
	printer.Test()

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	stacks := make(map[string]*awsCF.Stack)
	events := make(map[string][]*awsCF.StackEvent)

	stacks["foo-stack"] = &awsCF.Stack{
		StackId:     aws.String("foo-stack"),
		StackName:   aws.String("foo-stack"),
		StackStatus: aws.String("CREATE_COMPLETE"),
	}
	events["foo-stack"] = []*awsCF.StackEvent{}

	deleter := &MockStackDeleter{
		Stacks: stacks,
		Events: events,
	}

	assert.NotPanics(
		t,
		func() {
			taskDelete(
				deleter,
				objectStore,
				"test.yaml",   // templatePath
				"foo-stack", // stackName
				"ci",          // envName
				"profile",     // profile
				"region",      // region
			)
		},
	)
}

func TestDeleteTaskStackNotFound(t *testing.T) {
	printer.Test()

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	deleter := &MockStackDeleter{}

	assert.Panics(
		t,
		func() {
			taskDelete(
				deleter,
				objectStore,
				"test.yaml",   // templatePath
				"event-stack", // stackName
				"ci",          // envName
				"profile",     // profile
				"region",      // region
			)
		},
	)
}
