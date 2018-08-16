package tasks

import (
	"fmt"
	"os"
	"time"

	printer "github.com/KablamoOSS/go-cli-printer"

	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

func UpsertStackBody(
	templateBody []byte,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName string,
	cf *awsCF.CloudFormation,
	tags map[string]string,
) {
	changeSetIn := &awsCF.CreateChangeSetInput{
		Capabilities:  capabilities,
		ChangeSetName: aws.String(fmt.Sprintf("%s-upsert", stackName)),
		Description:   aws.String(fmt.Sprintf("Kombustion upsert of %s", stackName)),
		Parameters:    parameters,
		StackName:     aws.String(stackName),
		Tags:          formatTags(tags),
		TemplateBody:  aws.String(string(templateBody)),
	}
	upsertStack(cf, changeSetIn)
}

func UpsertStackURL(
	templateURL string,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName string,
	cf *awsCF.CloudFormation,
	tags map[string]string,
) {
	changeSetIn := &awsCF.CreateChangeSetInput{
		Capabilities:  capabilities,
		ChangeSetName: aws.String(fmt.Sprintf("%s-upsert", stackName)),
		Description:   aws.String(fmt.Sprintf("Kombustion upsert of %s", stackName)),
		Parameters:    parameters,
		StackName:     aws.String(stackName),
		Tags:          formatTags(tags),
		TemplateURL:   aws.String(templateURL),
	}
	upsertStack(cf, changeSetIn)
}

func upsertStack(
	cf *awsCF.CloudFormation,
	changeSetIn *awsCF.CreateChangeSetInput,
) {

	var err error
	var action string

	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: changeSetIn.StackName})
	if err == nil { //update
		action = "Updating"
		changeSetIn.ChangeSetType = aws.String("UPDATE")
	} else {
		action = "Creating"
		changeSetIn.ChangeSetType = aws.String("CREATE")
	}

	printer.Step("Creating change set")
	changeSetOut, err := cf.CreateChangeSet(changeSetIn)
	checkError(err)

	printer.Step("Waiting for change set creation")
	describeChangeSetIn := &awsCF.DescribeChangeSetInput{
		ChangeSetName: changeSetOut.Id,
	}
	cf.WaitUntilChangeSetCreateComplete(describeChangeSetIn)

	changeSet, err := cf.DescribeChangeSet(describeChangeSetIn)
	checkError(err)

	if *changeSet.Status == "FAILED" {
		printer.Fatal(
			fmt.Errorf("Cloudformation ChangeSet failed to create"),
			*changeSet.StatusReason,
			"",
		)
	}

	// TODO: In theory, a DescribeChangeSetOutput can be paginated (indicated
	// by having a .Next token), which may occur on large enough templates
	// (total response body > 1MB). Seems unlikely for most cases, but we
	// should probably handle it properly.

	printer.SubStep("Changes to be applied:", 1, true, true)
	for _, change := range changeSet.Changes {
		resChange := change.ResourceChange
		// TODO: ResourceChange has a .Replacement field to indicate whether a
		// Modify action will update in place or recreate the resource. We
		// should probably stitch that into the output.
		printer.SubStep(
			fmt.Sprintf(
				"%s %s %s",
				*resChange.Action,
				*resChange.LogicalResourceId,
				*resChange.ResourceType,
			),
			1,
			true,
			true,
		)
	}

	// TODO: Since we're displaying the action CF is going to take, we could
	// prompt the user to ask if they want to proceed.

	printer.Step("Executing change set")
	executeCSIn := &awsCF.ExecuteChangeSetInput{
		ChangeSetName: changeSetOut.Id,
	}
	_, err = cf.ExecuteChangeSet(executeCSIn)
	checkError(err)

	processUpsert(*changeSetIn.StackName, action, cf)
}

func processUpsert(stackName, action string, cf *awsCF.CloudFormation) {
	var err error
	var status *awsCF.DescribeStacksOutput

	PrintStackEventHeader()
	for {
		printer.Progress(action)
		time.Sleep(2 * time.Second)
		status, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkError(err)

		events, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
		checkError(err)

		if len(status.Stacks) > 0 {

			stack := status.Stacks[0]
			stackStatus := *stack.StackStatus

			if len(events.StackEvents) > 0 {
				PrintStackEvent(events.StackEvents[0], false)
			}

			if stackStatus != awsCF.StackStatusCreateInProgress &&
				stackStatus != awsCF.StackStatusUpdateInProgress &&
				stackStatus != awsCF.StackStatusUpdateCompleteCleanupInProgress {
				if stackStatus == awsCF.StackStatusCreateComplete ||
					stackStatus == awsCF.StackStatusUpdateComplete {
					printer.SubStep(
						fmt.Sprintf("Success %s Stack %s", action, stackName),
						1,
						true,
						true,
					)
					os.Exit(0)
				} else {
					printer.SubStep(
						fmt.Sprintf("Fail %s Stack %s", action, stackName),
						1,
						true,
						true,
					)

					printer.Error(fmt.Errorf("Upsert Failed"), "", "")
					time.Sleep(2 * time.Second)
					os.Exit(1)
				}
			}
		}
	}
}

func formatTags(tags map[string]string) []*awsCF.Tag {
	cfTags := make([]*awsCF.Tag, 0)
	for key, value := range tags {
		// Since aws-sdk-go insists on using string pointers, pointers to the
		// loop variables will have their values changed.
		// Creating a copy of the key / value here means we don't end up with
		// all the array elements referencing the same variable (and thus
		// having the same value).
		k := key
		v := value
		cfTags = append(cfTags, &awsCF.Tag{Key: &k, Value: &v})
	}
	return cfTags
}
