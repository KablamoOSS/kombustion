package tasks

import (
	"fmt"
	"os"
	"time"
	printer "github.com/KablamoOSS/go-cli-printer"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// DeleteStack removes a cloudformation stack
func DeleteStack(stackName, profile, region string) {
	cf := GetCloudformationClient(profile, region)
	printer.Step(fmt.Sprintf("Delete Stack %s:", stackName))

	//See if the stack exists to begin with
	_, err := cf.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
	checkError(err)

	_, err = cf.DeleteStack(&cloudformation.DeleteStackInput{StackName: aws.String(stackName)})
	checkError(err)

	// status polling
	PrintStackEventHeader()

	for {
		printer.Progress("Deleting")
		time.Sleep(2 * time.Second)
		status, err := cf.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
		checkErrorDeletePoll(err)

		events, _ := cf.DescribeStackEvents(&cloudformation.DescribeStackEventsInput{StackName: aws.String(stackName)})

		if len(status.Stacks) > 0 {
			stackStatus := *status.Stacks[0].StackStatus

			if len(events.StackEvents) > 0 {
				PrintStackEvent(events.StackEvents[0], false)
			}
			if stackStatus == cloudformation.StackStatusDeleteInProgress {
				continue
			}
		}
		break
	}

	// Make sure delete worked
	_, err = cf.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
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
