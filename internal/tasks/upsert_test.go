package tasks

import (
	"fmt"
	"testing"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/coretest"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/stretchr/testify/assert"
)

type MockStackUpserter struct {
	AcctID  string
	Events  map[string][]*awsCF.StackEvent
	Stacks  map[string]*awsCF.Stack
	Changes map[string]*awsCF.ChangeSetSummary
}

func (msu *MockStackUpserter) Open(_, _ string) string {
	msu.Events = make(map[string][]*awsCF.StackEvent)
	msu.Stacks = make(map[string]*awsCF.Stack)
	msu.Changes = make(map[string]*awsCF.ChangeSetSummary)
	return msu.AcctID
}

func (msu *MockStackUpserter) DescribeStackEvents(input *awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error) {
	events, ok := msu.Events[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("stack %s not found", *input.StackName)
	}

	out := &awsCF.DescribeStackEventsOutput{
		StackEvents: events,
	}

	return out, nil
}

func (msu *MockStackUpserter) DescribeStackEventsPages(input *awsCF.DescribeStackEventsInput, fn func(*awsCF.DescribeStackEventsOutput, bool) bool) error {
	out, err := msu.DescribeStackEvents(input)
	if err != nil {
		return err
	}

	fn(out, true)
	return nil
}

func (msu *MockStackUpserter) DescribeStacks(input *awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error) {
	stack, ok := msu.Stacks[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("stack not found: %s", *input.StackName)
	}
	output := &awsCF.DescribeStacksOutput{
		Stacks: []*awsCF.Stack{stack},
	}
	switch *stack.StackStatus {
	case "CREATE_IN_PROGRESS":
		stack.StackStatus = aws.String("CREATE_COMPLETE")
	case "UPDATE_IN_PROGRESS":
		stack.StackStatus = aws.String("UPDATE_COMPLETE")
	}
	return output, nil
}

func (msu *MockStackUpserter) CreateChangeSet(input *awsCF.CreateChangeSetInput) (*awsCF.CreateChangeSetOutput, error) {
	stack, ok := msu.Stacks[*input.StackName]
	if !ok && *input.ChangeSetType != "CREATE" {
		return nil, fmt.Errorf("stack not found: %s", *input.StackName)
	} else if ok && *input.ChangeSetType == "CREATE" {
		return nil, fmt.Errorf("stack exists: %s", *input.StackName)
	}

	changeId := fmt.Sprintf("%s-change", *input.StackName)
	change := &awsCF.ChangeSetSummary{
		ChangeSetId:     aws.String(changeId),
		ChangeSetName:   aws.String(changeId),
		ExecutionStatus: aws.String("UNAVAILABLE"),
		StackName:       input.StackName,
	}

	if stack != nil {
		change.StackId = stack.StackId
		stack.StackStatus = aws.String("REVIEW_IN_PROGRESS")
	}

	msu.Changes[changeId] = change

	output := &awsCF.CreateChangeSetOutput{
		Id:      change.ChangeSetId,
		StackId: change.StackId,
	}

	return output, nil
}

func (msu *MockStackUpserter) DescribeChangeSet(input *awsCF.DescribeChangeSetInput) (*awsCF.DescribeChangeSetOutput, error) {
	change, ok := msu.Changes[*input.ChangeSetName]
	if !ok {
		return nil, fmt.Errorf("change not found: %s", *input.ChangeSetName)
	}

	output := &awsCF.DescribeChangeSetOutput{
		ChangeSetId:     change.ChangeSetId,
		ChangeSetName:   change.ChangeSetName,
		Changes:         []*awsCF.Change{},
		ExecutionStatus: change.ExecutionStatus,
		Status:          aws.String("OK"),
	}
	return output, nil
}

func (msu *MockStackUpserter) WaitUntilChangeSetCreateComplete(input *awsCF.DescribeChangeSetInput) error {
	change, ok := msu.Changes[*input.ChangeSetName]
	if !ok {
		return fmt.Errorf("change not found: %s", *input.ChangeSetName)
	}
	change.ExecutionStatus = aws.String("AVAILABLE")
	return nil
}

func (msu *MockStackUpserter) DeleteChangeSet(input *awsCF.DeleteChangeSetInput) (*awsCF.DeleteChangeSetOutput, error) {
	change, ok := msu.Changes[*input.ChangeSetName]
	if !ok {
		return nil, fmt.Errorf("change not found: %s", *input.ChangeSetName)
	}
	delete(msu.Changes, *change.ChangeSetId)
	return &awsCF.DeleteChangeSetOutput{}, nil
}

func (msu *MockStackUpserter) DeleteStack(input *awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error) {
	stack, ok := msu.Stacks[*input.StackName]
	if !ok {
		return nil, fmt.Errorf("change not found: %s", *input.StackName)
	}
	delete(msu.Stacks, *stack.StackName)
	return &awsCF.DeleteStackOutput{}, nil
}

func (msu *MockStackUpserter) ExecuteChangeSet(input *awsCF.ExecuteChangeSetInput) (*awsCF.ExecuteChangeSetOutput, error) {
	change, ok := msu.Changes[*input.ChangeSetName]
	if !ok {
		return nil, fmt.Errorf("change not found: %s", *input.ChangeSetName)
	}

	stack, ok := msu.Stacks[*change.StackName]
	if !ok {
		msu.Stacks[*change.StackName] = &awsCF.Stack{
			StackId:     change.StackName,
			StackName:   change.StackName,
			StackStatus: aws.String("CREATE_IN_PROGRESS"),
		}
		msu.Events[*change.StackName] = []*awsCF.StackEvent{}
	} else {
		stack.StackStatus = aws.String("UPDATE_IN_PROGRESS")
	}

	return &awsCF.ExecuteChangeSetOutput{}, nil
}

func TestUpsert(t *testing.T) {
	printer.Test()
	client := &MockStackUpserter{
		AcctID: "12345",
	}

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	assert.NotPanics(
		t,
		func() {
			upsert(
				client,
				objectStore,
				"test.yaml",         // fileName
				"foo-stack",         //stackName
				"profile",           // profile
				"region",            // region
				map[string]string{}, // paramsMap
				"",                  // inputParameters
				map[string]string{}, // tagsMap
				"",                  // devPluginPath
				"ci",                // envName
				false,               // generateDefaultOutputs
				[]*string{},         // capabilities
				false,               // confirm
				"kombustion.yaml",   // manifest location
			)
		},
	)
}
