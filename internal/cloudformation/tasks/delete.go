package tasks

import (
	"fmt"
	printer "github.com/KablamoOSS/go-cli-printer"
	"os"
	"time"

	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

// DeleteStack removes a cloudformation stack
func DeleteStack(client cloudformation.StackDeleter, stackName, region string) {
	printer.Step(fmt.Sprintf("Delete Stack %s:", stackName))

	//See if the stack exists to begin with
	_, err := client.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	checkError(err)

	_, err = client.DeleteStack(&awsCF.DeleteStackInput{StackName: aws.String(stackName)})
	checkError(err)

	// status polling
	PrintStackEventHeader()

	for {
		printer.Progress("Deleting")
		time.Sleep(2 * time.Second)
		status, err := client.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkErrorDeletePoll(err)

		events, err := client.DescribeStackEvents(
			&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)},
		)
		checkError(err)

		if len(status.Stacks) > 0 {
			stackStatus := *status.Stacks[0].StackStatus

			if len(events.StackEvents) > 0 {
				PrintStackEvent(events.StackEvents[0], false)
			}
			if stackStatus == awsCF.StackStatusDeleteInProgress {
				continue
			}
		}
		break
	}

	// Make sure delete worked
	_, err = client.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err != nil {
		checkErrorDeletePoll(err)
	} else {
		printer.SubStep(
			fmt.Sprintf("Success Delete Stack %s", stackName),
			1,
			true,
			true,
		)
		os.Exit(0)
	}
}
