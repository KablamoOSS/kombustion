package tasks

import (
	"fmt"
	"testing"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/coretest"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/stretchr/testify/assert"
)

type MockStackEventer struct {
	AcctID string
	Events map[string][]*awsCF.StackEvent
}

func (mse *MockStackEventer) Open(_, _ string) string {
	return mse.AcctID
}

func (mse *MockStackEventer) DescribeStackEvents(input *awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error) {
	events, ok := mse.Events[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("stack %s not found", *input.StackName)
	}

	out := &awsCF.DescribeStackEventsOutput{
		StackEvents: events,
	}

	return out, nil
}

func (mse *MockStackEventer) DescribeStackEventsPages(input *awsCF.DescribeStackEventsInput, fn func(*awsCF.DescribeStackEventsOutput, bool) bool) error {
	out, err := mse.DescribeStackEvents(input)
	if err != nil {
		return err
	}

	fn(out, true)
	return nil
}

func TestEventsTask(t *testing.T) {
	printer.Test()

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	eventer := &MockStackEventer{
		AcctID: "acct-12345",
		Events: map[string][]*awsCF.StackEvent{
			"event-stack": []*awsCF.StackEvent{},
		},
	}

	assert.NotPanics(
		t,
		func() {
			printEvents(
				objectStore,
				eventer,
				"test.yaml",   // templatePath
				"event-stack", // stackName
				"ci",          // envName
				"profile",     // profile
				"region",      // region
			)
		},
	)
}

func TestEventsTaskNotFound(t *testing.T) {
	printer.Test()

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	eventer := &MockStackEventer{}

	assert.Panics(
		t,
		func() {
			printEvents(
				objectStore,
				eventer,
				"test.yaml",   // templatePath
				"event-stack", // stackName
				"ci",          // envName
				"profile",     // profile
				"region",      // region
			)
		},
	)
}
