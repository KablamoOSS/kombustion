package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/aws/aws-sdk-go/aws"
)

// PrintStackEvents outputs the events of a stack
// TODO: add flags to allow printing all events, and default only to recent events
func PrintStackEvents(cf *awsCF.CloudFormation, stackName string) {
	status, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
	checkError(err)

	printer.Step(fmt.Sprintf("Events for %s:", stackName))
	printer.SubStep(
		fmt.Sprintf(" %-19v | %-22v | %-30v | %v | %v | \n", "Time", "Status", "Type", "LogicalID", "Status Reason"),
		1,
		false,
		true,
	)
	for _, event := range status.StackEvents {
		if event.Timestamp != nil {
			printer.SubStep(
				fmt.Sprintf(" %-19v |", event.Timestamp.Format("2006-01-2 15:04:05")),
				1,
				false,
				true,
			)
		}
		if event.ResourceStatus != nil {
			printer.SubStep(
				fmt.Sprintf(" %-22v |", *event.ResourceStatus),
				1,
				false,
				true,
			)
		}
		if event.ResourceType != nil {
			printer.SubStep(
				fmt.Sprintf(" %-30v |", *event.ResourceType),
				1,
				false,
				true,
			)
		}
		if event.LogicalResourceId != nil {
			printer.SubStep(
				fmt.Sprintf(" %v |", *event.LogicalResourceId),
				1,
				false,
				true,
			)
		}
		if event.ResourceStatusReason != nil {
			printer.SubStep(
				fmt.Sprintf(" %v |", *event.ResourceStatusReason),
				1,
				false,
				true,
			)
		}
	}
}
