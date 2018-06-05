package tasks

import (
	"fmt"

	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/aws/aws-sdk-go/aws"
)

// PrintStackEvents outputs the events of a stack
// TODO: add flags to allow printing all events, and default only to recent events
func PrintStackEvents(cf *awsCF.CloudFormation, stackName string) {
	status, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
	checkError(err)

	fmt.Println()
	fmt.Printf(" %-19v | %-22v | %-30v | %v | %v | \n", "Time", "Status", "Type", "LogicalID", "Status Reason")
	for _, event := range status.StackEvents {
		if event.Timestamp != nil {
			fmt.Printf(" %-19v |", event.Timestamp.Format("2006-01-2 15:04:05"))
		}
		if event.ResourceStatus != nil {
			fmt.Printf(" %-22v |", *event.ResourceStatus)
		}
		if event.ResourceType != nil {
			fmt.Printf(" %-30v |", *event.ResourceType)
		}
		if event.LogicalResourceId != nil {
			fmt.Printf(" %v |", *event.LogicalResourceId)
		}
		if event.ResourceStatusReason != nil {
			fmt.Printf(" %v |", *event.ResourceStatusReason)
		}
		fmt.Println()
	}
}
