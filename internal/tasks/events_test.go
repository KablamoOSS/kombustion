package tasks

import (
	"fmt"
	"testing"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/coretest"
	"github.com/stretchr/testify/assert"
)

type MockStackEventer struct {
	AcctID string
	Events map[string][]*cloudformation.StackEvent
}

func (mse *MockStackEventer) Open(_, _ string) string {
	return mse.AcctID
}

func (mse *MockStackEventer) StackEvents(stackName string) ([]*cloudformation.StackEvent, error) {
	events, ok := mse.Events[stackName]
	if !ok {
		return nil, fmt.Errorf("stack %s not found", stackName)
	}
	return events, nil
}

func TestEventsTask(t *testing.T) {
	printer.Test()

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	eventer := &MockStackEventer{
		AcctID: "acct-12345",
		Events: map[string][]*cloudformation.StackEvent{
			"event-stack": []*cloudformation.StackEvent{},
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
